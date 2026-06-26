package models

const (
	ProviderMiniMax ModelProvider = "minimax"
)

const (
	MiniMaxM3          ModelID = "MiniMax-M3"
	MiniMaxM27         ModelID = "MiniMax-M2.7"
	MiniMaxM27Highspeed ModelID = "MiniMax-M2.7-highspeed"
	MiniMaxM25         ModelID = "MiniMax-M2.5"
	MiniMaxM25Highspeed ModelID = "MiniMax-M2.5-highspeed"
)

var MiniMaxModels = map[ModelID]Model{
	MiniMaxM3: {
		ID:                  MiniMaxM3,
		Name:                "MiniMax M3",
		Provider:            ProviderMiniMax,
		APIModel:            string(MiniMaxM3),
		CostPer1MIn:         0,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0,
		ContextWindow:        1_000_000,
		DefaultMaxTokens:     8192,
		CanReason:           true,
		SupportsAttachments:  true,
	},
	MiniMaxM27: {
		ID:                  MiniMaxM27,
		Name:                "MiniMax M2.7",
		Provider:            ProviderMiniMax,
		APIModel:            string(MiniMaxM27),
		CostPer1MIn:         0,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0,
		ContextWindow:        1_000_000,
		DefaultMaxTokens:     8192,
		CanReason:           true,
		SupportsAttachments:  true,
	},
	MiniMaxM27Highspeed: {
		ID:                  MiniMaxM27Highspeed,
		Name:                "MiniMax M2.7 Highspeed",
		Provider:            ProviderMiniMax,
		APIModel:            string(MiniMaxM27Highspeed),
		CostPer1MIn:         0,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0,
		ContextWindow:        1_000_000,
		DefaultMaxTokens:     8192,
		CanReason:           true,
		SupportsAttachments:  true,
	},
	MiniMaxM25: {
		ID:                  MiniMaxM25,
		Name:                "MiniMax M2.5",
		Provider:            ProviderMiniMax,
		APIModel:            string(MiniMaxM25),
		CostPer1MIn:         0,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0,
		ContextWindow:        1_000_000,
		DefaultMaxTokens:     8192,
		CanReason:           true,
		SupportsAttachments:  true,
	},
	MiniMaxM25Highspeed: {
		ID:                  MiniMaxM25Highspeed,
		Name:                "MiniMax M2.5 Highspeed",
		Provider:            ProviderMiniMax,
		APIModel:            string(MiniMaxM25Highspeed),
		CostPer1MIn:         0,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0,
		ContextWindow:        1_000_000,
		DefaultMaxTokens:     8192,
		CanReason:           true,
		SupportsAttachments:  true,
	},
}
