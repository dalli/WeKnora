package chat

import (
	"context"
	"fmt"

	"github.com/Tencent/WeKnora/internal/logger"
	"github.com/Tencent/WeKnora/internal/models/utils/lmstudio"
	"github.com/Tencent/WeKnora/internal/types"
	"github.com/sashabaranov/go-openai"
)

// LMStudioChat 实现了基于 LM Studio 的聊天
type LMStudioChat struct {
	modelName      string
	modelID        string
	lmStudioService *lmstudio.LMStudioService
	client         *openai.Client
}

// NewLMStudioChat 创建 LM Studio 聊天实例
func NewLMStudioChat(config *ChatConfig, lmStudioService *lmstudio.LMStudioService) (*LMStudioChat, error) {
	client := lmStudioService.GetClient()
	return &LMStudioChat{
		modelName:       config.ModelName,
		modelID:         config.ModelID,
		lmStudioService: lmStudioService,
		client:          client,
	}, nil
}

// convertMessages 转换消息格式为OpenAI格式
func (c *LMStudioChat) convertMessages(messages []Message) []openai.ChatCompletionMessage {
	openaiMessages := make([]openai.ChatCompletionMessage, 0, len(messages))
	for _, msg := range messages {
		openaiMsg := openai.ChatCompletionMessage{
			Role: msg.Role,
		}

		// 处理内容：对于 assistant 角色，内容可能为空（当有 tool_calls 时）
		if msg.Content != "" {
			openaiMsg.Content = msg.Content
		}

		// 处理 tool calls（assistant 角色）
		if len(msg.ToolCalls) > 0 {
			openaiMsg.ToolCalls = make([]openai.ToolCall, 0, len(msg.ToolCalls))
			for _, tc := range msg.ToolCalls {
				toolType := openai.ToolType(tc.Type)
				openaiMsg.ToolCalls = append(openaiMsg.ToolCalls, openai.ToolCall{
					ID:   tc.ID,
					Type: toolType,
					Function: openai.FunctionCall{
						Name:      tc.Function.Name,
						Arguments: tc.Function.Arguments,
					},
				})
			}
		}

		// 处理 tool 角色消息（工具返回结果）
		if msg.Role == "tool" {
			openaiMsg.ToolCallID = msg.ToolCallID
			openaiMsg.Name = msg.Name
		}

		openaiMessages = append(openaiMessages, openaiMsg)
	}
	return openaiMessages
}

// buildChatCompletionRequest 构建聊天请求参数
func (c *LMStudioChat) buildChatCompletionRequest(messages []Message, opts *ChatOptions, isStream bool) openai.ChatCompletionRequest {
	req := openai.ChatCompletionRequest{
		Model:    c.modelName,
		Messages: c.convertMessages(messages),
		Stream:   isStream,
	}

	// 添加可选参数
	if opts != nil {
		if opts.Temperature > 0 {
			req.Temperature = float32(opts.Temperature)
		}
		if opts.TopP > 0 {
			req.TopP = float32(opts.TopP)
		}
		if opts.MaxTokens > 0 {
			req.MaxTokens = opts.MaxTokens
		}
		if opts.MaxCompletionTokens > 0 {
			req.MaxCompletionTokens = opts.MaxCompletionTokens
		}
		if opts.FrequencyPenalty > 0 {
			req.FrequencyPenalty = float32(opts.FrequencyPenalty)
		}
		if opts.PresencePenalty > 0 {
			req.PresencePenalty = float32(opts.PresencePenalty)
		}

		// 处理 Tools（函数定义）
		if len(opts.Tools) > 0 {
			req.Tools = make([]openai.Tool, 0, len(opts.Tools))
			for _, tool := range opts.Tools {
				toolType := openai.ToolType(tool.Type)
				openaiTool := openai.Tool{
					Type: toolType,
					Function: &openai.FunctionDefinition{
						Name:        tool.Function.Name,
						Description: tool.Function.Description,
					},
				}
				// 转换 Parameters (map[string]interface{} -> JSON Schema)
				if tool.Function.Parameters != nil {
					openaiTool.Function.Parameters = tool.Function.Parameters
				}
				req.Tools = append(req.Tools, openaiTool)
			}
		}

		// 处理 ToolChoice
		if opts.ToolChoice != "" {
			switch opts.ToolChoice {
			case "none", "required", "auto":
				req.ToolChoice = opts.ToolChoice
			default:
				// 特定工具名称，使用 ToolChoice 对象
				req.ToolChoice = openai.ToolChoice{
					Type: "function",
					Function: openai.ToolFunction{
						Name: opts.ToolChoice,
					},
				}
			}
		}
	}

	return req
}

// Chat 进行非流式聊天
func (c *LMStudioChat) Chat(ctx context.Context, messages []Message, opts *ChatOptions) (*types.ChatResponse, error) {
	// 确保服务可用
	if err := c.ensureServiceAvailable(ctx); err != nil {
		return nil, err
	}

	// 构建请求参数
	req := c.buildChatCompletionRequest(messages, opts, false)

	// 记录请求日志
	logger.GetLogger(ctx).Infof("发送聊天请求到 LM Studio 模型 %s", c.modelName)

	// 发送请求
	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no response from LM Studio")
	}

	choice := resp.Choices[0]
	response := &types.ChatResponse{
		Content:      choice.Message.Content,
		FinishReason: string(choice.FinishReason),
		Usage: struct {
			PromptTokens     int `json:"prompt_tokens"`
			CompletionTokens int `json:"completion_tokens"`
			TotalTokens      int `json:"total_tokens"`
		}{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}

	// 转换 Tool Calls
	if len(choice.Message.ToolCalls) > 0 {
		response.ToolCalls = make([]types.LLMToolCall, 0, len(choice.Message.ToolCalls))
		for _, tc := range choice.Message.ToolCalls {
			response.ToolCalls = append(response.ToolCalls, types.LLMToolCall{
				ID:   tc.ID,
				Type: string(tc.Type),
				Function: types.FunctionCall{
					Name:      tc.Function.Name,
					Arguments: tc.Function.Arguments,
				},
			})
		}
	}

	return response, nil
}

// ChatStream 进行流式聊天
func (c *LMStudioChat) ChatStream(ctx context.Context, messages []Message, opts *ChatOptions) (<-chan types.StreamResponse, error) {
	// 确保服务可用
	if err := c.ensureServiceAvailable(ctx); err != nil {
		return nil, err
	}

	// 构建请求参数
	req := c.buildChatCompletionRequest(messages, opts, true)

	// 记录请求日志
	logger.GetLogger(ctx).Infof("发送流式聊天请求到 LM Studio 模型 %s", c.modelName)

	// 创建流式响应通道
	streamChan := make(chan types.StreamResponse)

	// 启动流式请求
	stream, err := c.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		close(streamChan)
		return nil, fmt.Errorf("create chat completion stream: %w", err)
	}

	// 在后台处理流式响应
	go func() {
		defer close(streamChan)
		defer stream.Close()

		toolCallMap := make(map[int]*types.LLMToolCall)
		lastFunctionName := make(map[int]string)
		nameNotified := make(map[int]bool)

		buildOrderedToolCalls := func() []types.LLMToolCall {
			if len(toolCallMap) == 0 {
				return nil
			}
			result := make([]types.LLMToolCall, 0, len(toolCallMap))
			for i := 0; i < len(toolCallMap); i++ {
				if tc, ok := toolCallMap[i]; ok && tc != nil {
					result = append(result, *tc)
				}
			}
			if len(result) == 0 {
				return nil
			}
			return result
		}

		for {
			response, err := stream.Recv()
			if err != nil {
				// 发送最后一个响应，包含收集到的 tool calls
				streamChan <- types.StreamResponse{
					ResponseType: types.ResponseTypeAnswer,
					Content:      "",
					Done:         true,
					ToolCalls:    buildOrderedToolCalls(),
				}
				return
			}

			if len(response.Choices) > 0 {
				delta := response.Choices[0].Delta
				isDone := string(response.Choices[0].FinishReason) != ""

				// 收集 tool calls（流式响应中 tool calls 可能分多次返回）
				if len(delta.ToolCalls) > 0 {
					for _, tc := range delta.ToolCalls {
						// 检查是否已经存在该 tool call（通过 index）
						var toolCallIndex int
						if tc.Index != nil {
							toolCallIndex = *tc.Index
						}
						toolCallEntry, exists := toolCallMap[toolCallIndex]
						if !exists || toolCallEntry == nil {
							toolCallEntry = &types.LLMToolCall{
								Type: string(tc.Type),
								Function: types.FunctionCall{
									Name:      "",
									Arguments: "",
								},
							}
							toolCallMap[toolCallIndex] = toolCallEntry
						}

						// 更新 ID、类型
						if tc.ID != "" {
							toolCallEntry.ID = tc.ID
						}
						if tc.Type != "" {
							toolCallEntry.Type = string(tc.Type)
						}

						// 累积函数名称（可能分多次返回）
						if tc.Function.Name != "" {
							toolCallEntry.Function.Name += tc.Function.Name
						}

						// 累积参数（可能为部分 JSON）
						argsUpdated := false
						if tc.Function.Arguments != "" {
							toolCallEntry.Function.Arguments += tc.Function.Arguments
							argsUpdated = true
						}

						currName := toolCallEntry.Function.Name
						if currName != "" &&
							currName == lastFunctionName[toolCallIndex] &&
							argsUpdated &&
							!nameNotified[toolCallIndex] &&
							toolCallEntry.ID != "" {
							streamChan <- types.StreamResponse{
								ResponseType: types.ResponseTypeToolCall,
								Content:      "",
								Done:         false,
								Data: map[string]interface{}{
									"tool_name":    currName,
									"tool_call_id": toolCallEntry.ID,
								},
							}
							nameNotified[toolCallIndex] = true
						}

						lastFunctionName[toolCallIndex] = currName
					}
				}

				// 发送内容块
				if delta.Content != "" {
					streamChan <- types.StreamResponse{
						ResponseType: types.ResponseTypeAnswer,
						Content:      delta.Content,
						Done:         isDone,
						ToolCalls:    buildOrderedToolCalls(),
					}
				}

				// 如果是最后一次响应，确保发送包含所有 tool calls 的响应
				if isDone && len(toolCallMap) > 0 {
					streamChan <- types.StreamResponse{
						ResponseType: types.ResponseTypeAnswer,
						Content:      "",
						Done:         true,
						ToolCalls:    buildOrderedToolCalls(),
					}
				}
			}
		}
	}()

	return streamChan, nil
}

// 确保服务可用
func (c *LMStudioChat) ensureServiceAvailable(ctx context.Context) error {
	logger.GetLogger(ctx).Infof("确保 LM Studio 服务可用")
	return c.lmStudioService.StartService(ctx)
}

// GetModelName 获取模型名称
func (c *LMStudioChat) GetModelName() string {
	return c.modelName
}

// GetModelID 获取模型ID
func (c *LMStudioChat) GetModelID() string {
	return c.modelID
}

