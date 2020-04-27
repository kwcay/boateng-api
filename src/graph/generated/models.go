// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"fmt"
	"io"
	"strconv"
)

type Alphabet struct {
	Code       string       `json:"code"`
	Name       string       `json:"name"`
	Script     *Script      `json:"script"`
	Characters *string      `json:"characters"`
	Languages  []*Language  `json:"languages"`
	References []*Reference `json:"references"`
}

type Expression struct {
	ID                   string             `json:"id"`
	Type                 ExpressionType     `json:"type"`
	Titles               []*Transliteration `json:"titles"`
	Languages            []*Language        `json:"languages"`
	PartOfSpeech         *PartOfSpeech      `json:"partOfSpeech"`
	NounType             *NounType          `json:"nounType"`
	Lexeme               *Expression        `json:"lexeme"`
	LiteralTranslation   *string            `json:"literalTranslation"`
	PracticalTranslation *string            `json:"practicalTranslation"`
	Meaning              *string            `json:"meaning"`
	Tags                 []*Tag             `json:"tags"`
	RelatedExpressions   []*Expression      `json:"relatedExpressions"`
	References           []*Reference       `json:"references"`
}

type Language struct {
	Code             string             `json:"code"`
	Names            []*Transliteration `json:"names"`
	Parent           *Language          `json:"parent"`
	Lexifier         *Language          `json:"lexifier"`
	RelatedLanguages []*Language        `json:"relatedLanguages"`
	GlottologID      *string            `json:"glottologId"`
	Alphabets        []*Alphabet        `json:"alphabets"`
	IsFamily         *bool              `json:"isFamily"`
	References       []*Reference       `json:"references"`
}

type Reference struct {
	Type ReferenceType `json:"type"`
	Mla  *string       `json:"mla"`
}

type Script struct {
	Code  string             `json:"code"`
	Names []*Transliteration `json:"names"`
}

type Story struct {
	Type     StoryType    `json:"type"`
	Lines    []*StoryLine `json:"lines"`
	Language *Language    `json:"language"`
	Script   *Script      `json:"script"`
}

type StoryLine struct {
	Story   *Story `json:"story"`
	Content string `json:"content"`
}

type Tag struct {
	Name string `json:"name"`
}

type Transliteration struct {
	Value                     string  `json:"value"`
	TransliterationLangCode   *string `json:"transliterationLangCode"`
	TransliterationScriptCode *string `json:"transliterationScriptCode"`
}

type ExpressionType string

const (
	ExpressionTypeExpression ExpressionType = "Expression"
	ExpressionTypePhrase     ExpressionType = "Phrase"
	ExpressionTypeProverb    ExpressionType = "Proverb"
	ExpressionTypeWord       ExpressionType = "Word"
)

var AllExpressionType = []ExpressionType{
	ExpressionTypeExpression,
	ExpressionTypePhrase,
	ExpressionTypeProverb,
	ExpressionTypeWord,
}

func (e ExpressionType) IsValid() bool {
	switch e {
	case ExpressionTypeExpression, ExpressionTypePhrase, ExpressionTypeProverb, ExpressionTypeWord:
		return true
	}
	return false
}

func (e ExpressionType) String() string {
	return string(e)
}

func (e *ExpressionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ExpressionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ExpressionType", str)
	}
	return nil
}

func (e ExpressionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NounType string

const (
	NounTypeName   NounType = "Name"
	NounTypePlace  NounType = "Place"
	NounTypePerson NounType = "Person"
)

var AllNounType = []NounType{
	NounTypeName,
	NounTypePlace,
	NounTypePerson,
}

func (e NounType) IsValid() bool {
	switch e {
	case NounTypeName, NounTypePlace, NounTypePerson:
		return true
	}
	return false
}

func (e NounType) String() string {
	return string(e)
}

func (e *NounType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NounType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NounType", str)
	}
	return nil
}

func (e NounType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PartOfSpeech string

const (
	PartOfSpeechAdjective    PartOfSpeech = "Adjective"
	PartOfSpeechAdverb       PartOfSpeech = "Adverb"
	PartOfSpeechConjunction  PartOfSpeech = "Conjunction"
	PartOfSpeechInterjection PartOfSpeech = "Interjection"
	PartOfSpeechNoun         PartOfSpeech = "Noun"
	PartOfSpeechPreposition  PartOfSpeech = "Preposition"
	PartOfSpeechPronoun      PartOfSpeech = "Pronoun"
	PartOfSpeechVerb         PartOfSpeech = "Verb"
	PartOfSpeechPrefix       PartOfSpeech = "Prefix"
	PartOfSpeechSuffix       PartOfSpeech = "Suffix"
)

var AllPartOfSpeech = []PartOfSpeech{
	PartOfSpeechAdjective,
	PartOfSpeechAdverb,
	PartOfSpeechConjunction,
	PartOfSpeechInterjection,
	PartOfSpeechNoun,
	PartOfSpeechPreposition,
	PartOfSpeechPronoun,
	PartOfSpeechVerb,
	PartOfSpeechPrefix,
	PartOfSpeechSuffix,
}

func (e PartOfSpeech) IsValid() bool {
	switch e {
	case PartOfSpeechAdjective, PartOfSpeechAdverb, PartOfSpeechConjunction, PartOfSpeechInterjection, PartOfSpeechNoun, PartOfSpeechPreposition, PartOfSpeechPronoun, PartOfSpeechVerb, PartOfSpeechPrefix, PartOfSpeechSuffix:
		return true
	}
	return false
}

func (e PartOfSpeech) String() string {
	return string(e)
}

func (e *PartOfSpeech) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PartOfSpeech(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PartOfSpeech", str)
	}
	return nil
}

func (e PartOfSpeech) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ReferenceType string

const (
	ReferenceTypeArticle      ReferenceType = "Article"
	ReferenceTypeAudio        ReferenceType = "Audio"
	ReferenceTypeBook         ReferenceType = "Book"
	ReferenceTypeEncyclopedia ReferenceType = "Encyclopedia"
	ReferenceTypeFilm         ReferenceType = "Film"
	ReferenceTypeInterview    ReferenceType = "Interview"
	ReferenceTypePaper        ReferenceType = "Paper"
	ReferenceTypePerson       ReferenceType = "Person"
	ReferenceTypeReport       ReferenceType = "Report"
	ReferenceTypeSocialMedia  ReferenceType = "SocialMedia"
	ReferenceTypeSong         ReferenceType = "Song"
	ReferenceTypeVideo        ReferenceType = "Video"
	ReferenceTypeWebsite      ReferenceType = "Website"
	ReferenceTypeOther        ReferenceType = "Other"
)

var AllReferenceType = []ReferenceType{
	ReferenceTypeArticle,
	ReferenceTypeAudio,
	ReferenceTypeBook,
	ReferenceTypeEncyclopedia,
	ReferenceTypeFilm,
	ReferenceTypeInterview,
	ReferenceTypePaper,
	ReferenceTypePerson,
	ReferenceTypeReport,
	ReferenceTypeSocialMedia,
	ReferenceTypeSong,
	ReferenceTypeVideo,
	ReferenceTypeWebsite,
	ReferenceTypeOther,
}

func (e ReferenceType) IsValid() bool {
	switch e {
	case ReferenceTypeArticle, ReferenceTypeAudio, ReferenceTypeBook, ReferenceTypeEncyclopedia, ReferenceTypeFilm, ReferenceTypeInterview, ReferenceTypePaper, ReferenceTypePerson, ReferenceTypeReport, ReferenceTypeSocialMedia, ReferenceTypeSong, ReferenceTypeVideo, ReferenceTypeWebsite, ReferenceTypeOther:
		return true
	}
	return false
}

func (e ReferenceType) String() string {
	return string(e)
}

func (e *ReferenceType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReferenceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReferenceType", str)
	}
	return nil
}

func (e ReferenceType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StoryType string

const (
	StoryTypePoem  StoryType = "Poem"
	StoryTypeSong  StoryType = "Song"
	StoryTypeStory StoryType = "Story"
)

var AllStoryType = []StoryType{
	StoryTypePoem,
	StoryTypeSong,
	StoryTypeStory,
}

func (e StoryType) IsValid() bool {
	switch e {
	case StoryTypePoem, StoryTypeSong, StoryTypeStory:
		return true
	}
	return false
}

func (e StoryType) String() string {
	return string(e)
}

func (e *StoryType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StoryType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StoryType", str)
	}
	return nil
}

func (e StoryType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}