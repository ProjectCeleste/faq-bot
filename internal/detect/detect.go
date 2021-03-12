package detect

import (
	"regexp"
	"unicode"
)

// QuestionDetection type of message detection.
type QuestionDetection int

const (
	// QuestionBoth ignores if the message is a question or not
	QuestionBoth QuestionDetection = iota

	// QuestionOnly only detect if sentence ends with a question mark
	QuestionOnly

	// QuestionNo only detect if sentence doesn't end with a question mark
	QuestionNo
)

var (
	punctuation      = regexp.MustCompile(`[,.!?:;]`)
	nextQuestionMark = regexp.MustCompile(`^[^.!]*\?`)
)

// SentenceDetector criteria for specific sentence detection.
type SentenceDetector struct {
	Question QuestionDetection
	Groups   []WordGroup
	Variants []SentenceDetector
}

// WordGroup a group of words of the same type used for sentence detection.
type WordGroup struct {
	Negate bool
	Words  []string
}

// Detect if the given sentence matches the criteria of the SentenceDetector.
func (d *SentenceDetector) Detect(sentence string) bool {
	sentenceOK := true
	s := sentence
	for _, g := range d.Groups {
		i, reset, ok := g.Detect(s)
		if reset {
			if !d.Detect(s[i:]) {
				sentenceOK = false
			}
			break
		}
		if ok {
			s = s[i:]
		} else {
			sentenceOK = false
			break
		}
	}
	if sentenceOK && d.Question != QuestionBoth {
		questionMarkFound := nextQuestionMark.MatchString(s)
		switch d.Question {
		case QuestionOnly:
			sentenceOK = questionMarkFound
		case QuestionNo:
			sentenceOK = !questionMarkFound
		}
	}

	if !sentenceOK {
		for _, v := range d.Variants {
			if v.Detect(sentence) {
				return true
			}
		}
	}

	return sentenceOK
}

// Detect if the given list of words matches the WordGroup.
func (d *WordGroup) Detect(sentence string) (int, bool, bool) {
	wordBegin := -1
	newI := 0
	for i, c := range sentence {
		if !unicode.IsLetter(c) && c != '\'' {
			if wordBegin != -1 {
				newI = i
				if contains(d.Words, sentence[wordBegin:i]) {
					return newI, false, !d.Negate
				}
				wordBegin = -1
			}
			if isEndOfSentence(c) {
				if d.Negate {
					return 0, false, true
				}
				return i + 1, true, false
			}
			continue
		}
		if wordBegin == -1 {
			wordBegin = i
		}
	}

	// String ended with a letter
	if wordBegin != -1 && contains(d.Words, sentence[wordBegin:]) {
		return newI, false, !d.Negate
	}
	if d.Negate {
		return 0, false, true
	}

	return newI, false, false
}

func contains(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

func removePunctuation(sentence string) string {
	return punctuation.ReplaceAllString(sentence, "")
}

func isEndOfSentence(chr rune) bool {
	return chr == '!' || chr == '?' || chr == '.'
}
