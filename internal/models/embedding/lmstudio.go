package embedding

import (
	"context"
	"fmt"
	"time"

	"github.com/Tencent/WeKnora/internal/logger"
	"github.com/Tencent/WeKnora/internal/models/utils/lmstudio"
	"github.com/sashabaranov/go-openai"
)

// LMStudioEmbedder implements text vectorization functionality using LM Studio
type LMStudioEmbedder struct {
	modelName            string
	truncatePromptTokens int
	lmStudioService      *lmstudio.LMStudioService
	client               *openai.Client
	dimensions           int
	modelID              string
	EmbedderPooler
}

// NewLMStudioEmbedder creates a new LM Studio embedder
func NewLMStudioEmbedder(
	baseURL,
	modelName string,
	truncatePromptTokens int,
	dimensions int,
	modelID string,
	pooler EmbedderPooler,
	lmStudioService *lmstudio.LMStudioService,
) (*LMStudioEmbedder, error) {
	if modelName == "" {
		return nil, fmt.Errorf("model name is required")
	}

	if truncatePromptTokens == 0 {
		truncatePromptTokens = 511
	}

	return &LMStudioEmbedder{
		modelName:            modelName,
		truncatePromptTokens: truncatePromptTokens,
		lmStudioService:      lmStudioService,
		client:               lmStudioService.GetClient(),
		EmbedderPooler:       pooler,
		dimensions:           dimensions,
		modelID:              modelID,
	}, nil
}

// ensureServiceAvailable ensures that the service is available
func (e *LMStudioEmbedder) ensureServiceAvailable(ctx context.Context) error {
	logger.GetLogger(ctx).Infof("Ensuring LM Studio service is available")
	return e.lmStudioService.StartService(ctx)
}

// Embed converts text to vector
func (e *LMStudioEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {
	embedding, err := e.BatchEmbed(ctx, []string{text})
	if err != nil {
		return nil, fmt.Errorf("failed to embed text: %w", err)
	}

	if len(embedding) == 0 {
		return nil, fmt.Errorf("failed to embed text: no embedding returned")
	}

	return embedding[0], nil
}

// BatchEmbed converts multiple texts to vectors in batch
func (e *LMStudioEmbedder) BatchEmbed(ctx context.Context, texts []string) ([][]float32, error) {
	// Ensure service is available
	if err := e.ensureServiceAvailable(ctx); err != nil {
		return nil, err
	}

	// Create request
	req := openai.EmbeddingRequest{
		Model: openai.EmbeddingModel(e.modelName),
		Input: texts,
	}

	// Send request
	startTime := time.Now()
	resp, err := e.client.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get embedding vectors: %w", err)
	}

	logger.GetLogger(ctx).Debugf("Embedding vector retrieval took: %v", time.Since(startTime))

	// Extract embeddings
	embeddings := make([][]float32, len(resp.Data))
	for i, data := range resp.Data {
		embeddings[i] = data.Embedding
	}

	return embeddings, nil
}

// GetModelName returns the model name
func (e *LMStudioEmbedder) GetModelName() string {
	return e.modelName
}

// GetDimensions returns the vector dimensions
func (e *LMStudioEmbedder) GetDimensions() int {
	return e.dimensions
}

// GetModelID returns the model ID
func (e *LMStudioEmbedder) GetModelID() string {
	return e.modelID
}

