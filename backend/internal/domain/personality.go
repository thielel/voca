package domain

import "time"

// Trait represents the Big Five personality traits
type Trait string

const (
	TraitExtraversion       Trait = "extraversion"
	TraitAgreeableness      Trait = "agreeableness"
	TraitConscientiousness  Trait = "conscientiousness"
	TraitEmotionalStability Trait = "emotional_stability"
	TraitOpenness           Trait = "openness"
)

// Question represents an IPIP Big Five questionnaire item
type Question struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Trait    Trait  `json:"trait"`
	Reversed bool   `json:"reversed"`
}

// Answer represents a user's response to a question
type Answer struct {
	QuestionID int `json:"question_id"`
	Value      int `json:"value"` // 1-5 Likert scale
}

// PersonalityResult stores the calculated personality scores
type PersonalityResult struct {
	ID                 string            `json:"id"`
	SessionID          string            `json:"session_id"`
	Extraversion       float64           `json:"extraversion"`
	Agreeableness      float64           `json:"agreeableness"`
	Conscientiousness  float64           `json:"conscientiousness"`
	EmotionalStability float64           `json:"emotional_stability"`
	Openness           float64           `json:"openness"`
	CreatedAt          time.Time         `json:"created_at"`
	Interpretations    map[Trait]string  `json:"interpretations,omitempty"`
}

// TraitInterpretation stores an AI-generated interpretation for a specific trait
type TraitInterpretation struct {
	ID             string    `json:"id"`
	ResultID       string    `json:"result_id"`
	Trait          Trait     `json:"trait"`
	Interpretation string    `json:"interpretation"`
	CreatedAt      time.Time `json:"created_at"`
}

// TraitDisplayName returns the German display name for a trait
func (t Trait) DisplayName() string {
	switch t {
	case TraitExtraversion:
		return "Extraversion"
	case TraitAgreeableness:
		return "Verträglichkeit"
	case TraitConscientiousness:
		return "Gewissenhaftigkeit"
	case TraitEmotionalStability:
		return "Emotionale Stabilität"
	case TraitOpenness:
		return "Offenheit für Erfahrungen"
	default:
		return string(t)
	}
}

// SubmitAnswersRequest is the request body for submitting questionnaire answers
type SubmitAnswersRequest struct {
	SessionID string   `json:"session_id"`
	Answers   []Answer `json:"answers"`
	Language  string   `json:"language,omitempty"` // Optional: language for AI interpretations (defaults to "de")
}

// SubmitAnswersResponse is the response after calculating results
type SubmitAnswersResponse struct {
	Result PersonalityResult `json:"result"`
}

// GetQuestionsResponse returns all questionnaire items
type GetQuestionsResponse struct {
	Questions []Question `json:"questions"`
}

// GetQuestions returns the IPIP Big Five 50-item questionnaire
// Source: International Personality Item Pool (https://ipip.ori.org/)
// License: Public Domain
func GetQuestions() []Question {
	return []Question{
		// Items 1-10: First round of all 5 traits (alternating polarity)
		{ID: 1, Text: "Am the life of the party.", Trait: TraitExtraversion, Reversed: false},
		{ID: 2, Text: "Feel little concern for others.", Trait: TraitAgreeableness, Reversed: true},
		{ID: 3, Text: "Am always prepared.", Trait: TraitConscientiousness, Reversed: false},
		{ID: 4, Text: "Get stressed out easily.", Trait: TraitEmotionalStability, Reversed: true},
		{ID: 5, Text: "Have a rich vocabulary.", Trait: TraitOpenness, Reversed: false},
		{ID: 6, Text: "Don't talk a lot.", Trait: TraitExtraversion, Reversed: true},
		{ID: 7, Text: "Am interested in people.", Trait: TraitAgreeableness, Reversed: false},
		{ID: 8, Text: "Leave my belongings around.", Trait: TraitConscientiousness, Reversed: true},
		{ID: 9, Text: "Am relaxed most of the time.", Trait: TraitEmotionalStability, Reversed: false},
		{ID: 10, Text: "Have difficulty understanding abstract ideas.", Trait: TraitOpenness, Reversed: true},

		// Items 11-20: Second round
		{ID: 11, Text: "Feel comfortable around people.", Trait: TraitExtraversion, Reversed: false},
		{ID: 12, Text: "Insult people.", Trait: TraitAgreeableness, Reversed: true},
		{ID: 13, Text: "Pay attention to details.", Trait: TraitConscientiousness, Reversed: false},
		{ID: 14, Text: "Worry about things.", Trait: TraitEmotionalStability, Reversed: true},
		{ID: 15, Text: "Have a vivid imagination.", Trait: TraitOpenness, Reversed: false},
		{ID: 16, Text: "Keep in the background.", Trait: TraitExtraversion, Reversed: true},
		{ID: 17, Text: "Sympathize with others' feelings.", Trait: TraitAgreeableness, Reversed: false},
		{ID: 18, Text: "Make a mess of things.", Trait: TraitConscientiousness, Reversed: true},
		{ID: 19, Text: "Seldom feel blue.", Trait: TraitEmotionalStability, Reversed: false},
		{ID: 20, Text: "Am not interested in abstract ideas.", Trait: TraitOpenness, Reversed: true},

		// Items 21-30: Third round
		{ID: 21, Text: "Start conversations.", Trait: TraitExtraversion, Reversed: false},
		{ID: 22, Text: "Am not interested in other people's problems.", Trait: TraitAgreeableness, Reversed: true},
		{ID: 23, Text: "Get chores done right away.", Trait: TraitConscientiousness, Reversed: false},
		{ID: 24, Text: "Am easily disturbed.", Trait: TraitEmotionalStability, Reversed: true},
		{ID: 25, Text: "Have excellent ideas.", Trait: TraitOpenness, Reversed: false},
		{ID: 26, Text: "Have little to say.", Trait: TraitExtraversion, Reversed: true},
		{ID: 27, Text: "Have a soft heart.", Trait: TraitAgreeableness, Reversed: false},
		{ID: 28, Text: "Often forget to put things back in their proper place.", Trait: TraitConscientiousness, Reversed: true},
		{ID: 29, Text: "Get upset easily.", Trait: TraitEmotionalStability, Reversed: true},
		{ID: 30, Text: "Do not have a good imagination.", Trait: TraitOpenness, Reversed: true},

		// Items 31-40: Fourth round
		{ID: 31, Text: "Talk to a lot of different people at parties.", Trait: TraitExtraversion, Reversed: false},
		{ID: 32, Text: "Am not really interested in others.", Trait: TraitAgreeableness, Reversed: true},
		{ID: 33, Text: "Like order.", Trait: TraitConscientiousness, Reversed: false},
		{ID: 34, Text: "Change my mood a lot.", Trait: TraitEmotionalStability, Reversed: true},
		{ID: 35, Text: "Am quick to understand things.", Trait: TraitOpenness, Reversed: false},
		{ID: 36, Text: "Don't like to draw attention to myself.", Trait: TraitExtraversion, Reversed: true},
		{ID: 37, Text: "Take time out for others.", Trait: TraitAgreeableness, Reversed: false},
		{ID: 38, Text: "Shirk my duties.", Trait: TraitConscientiousness, Reversed: true},
		{ID: 39, Text: "Have frequent mood swings.", Trait: TraitEmotionalStability, Reversed: true},
		{ID: 40, Text: "Use difficult words.", Trait: TraitOpenness, Reversed: false},

		// Items 41-50: Fifth round
		{ID: 41, Text: "Don't mind being the center of attention.", Trait: TraitExtraversion, Reversed: false},
		{ID: 42, Text: "Feel others' emotions.", Trait: TraitAgreeableness, Reversed: false},
		{ID: 43, Text: "Follow a schedule.", Trait: TraitConscientiousness, Reversed: false},
		{ID: 44, Text: "Get irritated easily.", Trait: TraitEmotionalStability, Reversed: true},
		{ID: 45, Text: "Spend time reflecting on things.", Trait: TraitOpenness, Reversed: false},
		{ID: 46, Text: "Am quiet around strangers.", Trait: TraitExtraversion, Reversed: true},
		{ID: 47, Text: "Make people feel at ease.", Trait: TraitAgreeableness, Reversed: false},
		{ID: 48, Text: "Am exacting in my work.", Trait: TraitConscientiousness, Reversed: false},
		{ID: 49, Text: "Often feel blue.", Trait: TraitEmotionalStability, Reversed: true},
		{ID: 50, Text: "Am full of ideas.", Trait: TraitOpenness, Reversed: false},
	}
}
