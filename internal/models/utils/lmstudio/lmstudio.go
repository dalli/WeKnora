package lmstudio

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/Tencent/WeKnora/internal/logger"
	"github.com/sashabaranov/go-openai"
)

// LMStudioService manages LM Studio service
type LMStudioService struct {
	client      *openai.Client
	baseURL     string
	mu          sync.Mutex
	isAvailable bool
	isOptional  bool // marks if LM Studio service is optional
}

// LMStudioModelInfo represents detailed information about an LM Studio model
type LMStudioModelInfo struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	OwnedBy string `json:"owned_by"`
}

// GetLMStudioService gets LM Studio service instance (singleton pattern)
func GetLMStudioService() (*LMStudioService, error) {
	// Get LM Studio base URL from environment variable, if not set use default value
	logger.GetLogger(context.Background()).Infof("LM Studio base URL: %s", os.Getenv("LM_STUDIO_BASE_URL"))
	baseURL := "http://localhost:1234/v1"
	envURL := os.Getenv("LM_STUDIO_BASE_URL")
	if envURL != "" {
		baseURL = envURL
	}

	// Ensure base URL ends with /v1
	if baseURL[len(baseURL)-3:] != "/v1" {
		if baseURL[len(baseURL)-1] != '/' {
			baseURL += "/v1"
		} else {
			baseURL += "v1"
		}
	}

	// Create URL object
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid LM Studio service URL: %w", err)
	}

	// Create OpenAI client with LM Studio base URL
	// LM Studio doesn't require API key, but go-openai requires one, so we use a dummy key
	config := openai.DefaultConfig("lm-studio")
	config.BaseURL = parsedURL.String()

	// Check if LM Studio is set as optional
	isOptional := false
	if os.Getenv("LM_STUDIO_OPTIONAL") == "true" {
		isOptional = true
		logger.GetLogger(context.Background()).Info("LM Studio service set to optional mode")
	}

	service := &LMStudioService{
		client:     openai.NewClientWithConfig(config),
		baseURL:    baseURL,
		isOptional: isOptional,
	}

	return service, nil
}

// StartService checks if LM Studio service is available
func (s *LMStudioService) StartService(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if service is available by listing models
	// This is a lightweight operation that verifies the service is running
	_, err := s.client.ListModels(ctx)
	if err != nil {
		logger.GetLogger(ctx).Warnf("lm studio service unavailable: %v", err)
		s.isAvailable = false

		// If configured as optional, don't return an error
		if s.isOptional {
			logger.GetLogger(ctx).Info("lm studio service set as optional, will continue running the application")
			return nil
		}

		return fmt.Errorf("lm studio service unavailable: %w", err)
	}

	s.isAvailable = true
	return nil
}

// IsAvailable returns whether the service is available
func (s *LMStudioService) IsAvailable() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.isAvailable
}

// IsModelAvailable checks if a model is available
func (s *LMStudioService) IsModelAvailable(ctx context.Context, modelName string) (bool, error) {
	// First check if the service is available
	if err := s.StartService(ctx); err != nil {
		return false, err
	}

	// If service is not available but set as optional, return false but no error
	if !s.isAvailable && s.isOptional {
		return false, nil
	}

	// Get model list
	models, err := s.ListModels(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get model list: %w", err)
	}

	// Check if model is in the list
	for _, model := range models {
		if model.ID == modelName {
			return true, nil
		}
	}

	return false, nil
}

// ListModels lists all available models
func (s *LMStudioService) ListModels(ctx context.Context) ([]LMStudioModelInfo, error) {
	// First check if service is available
	if err := s.StartService(ctx); err != nil {
		return nil, err
	}

	// If service is not available but set as optional, return empty list
	if !s.IsAvailable() && s.isOptional {
		return []LMStudioModelInfo{}, nil
	}

	// Get model list from OpenAI API
	openaiModels, err := s.client.ListModels(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get model list: %w", err)
	}

	// Convert to LMStudioModelInfo
	models := make([]LMStudioModelInfo, len(openaiModels.Models))
	for i, model := range openaiModels.Models {
		models[i] = LMStudioModelInfo{
			ID:      model.ID,
			Object:  model.Object,
			Created: 0, // Created field may not be available in go-openai Model struct
			OwnedBy: model.OwnedBy,
		}
	}

	return models, nil
}

// ListModelNames lists all available model names only
func (s *LMStudioService) ListModelNames(ctx context.Context) ([]string, error) {
	models, err := s.ListModels(ctx)
	if err != nil {
		return nil, err
	}

	modelNames := make([]string, len(models))
	for i, model := range models {
		modelNames[i] = model.ID
	}

	return modelNames, nil
}

// GetClient returns the underlying OpenAI client for advanced operations
func (s *LMStudioService) GetClient() *openai.Client {
	return s.client
}

// GetBaseURL returns the base URL
func (s *LMStudioService) GetBaseURL() string {
	return s.baseURL
}

