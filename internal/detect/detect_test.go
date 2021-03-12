package detect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemovePunctuation(t *testing.T) {
	assert.Equal(t, "hello world", removePunctuation(".hell;:o, wor?ld!"))
}

func TestDetectSentence(t *testing.T) {
	detector := &SentenceDetector{
		Question: QuestionOnly,
		Groups: []WordGroup{
			{Words: []string{"when", "are", "have", "has"}},
			{Words: []string{"romans", "roman"}},
			{Words: []string{"out", "ready", "released", "release"}},
		},
		Variants: []SentenceDetector{
			{
				Question: QuestionOnly,
				Groups: []WordGroup{
					{Words: []string{"when"}},
					{Words: []string{"play", "access"}},
					{Words: []string{"romans", "roman"}},
				},
			},
			{
				Question: QuestionOnly,
				Groups: []WordGroup{
					{Words: []string{"romans", "roman"}},
					{Words: []string{"when", "out"}},
				},
			},
		},
	}
	assert.False(t, detector.Detect(""))
	assert.True(t, detector.Detect("when will the romans come out?"))
	assert.True(t, detector.Detect("when will the romans come out   ?"))
	assert.False(t, detector.Detect("when will the ro  mans come out?"))
	assert.False(t, detector.Detect("when will the romans come out"))
	assert.True(t, detector.Detect("when can i play the romans?"))
	assert.True(t, detector.Detect("are the romans ready?"))
	assert.True(t, detector.Detect("are the romans out?"))
	assert.True(t, detector.Detect("are the romans released?"))
	assert.True(t, detector.Detect("when will the romans release?"))
	assert.True(t, detector.Detect("are the romans out yet?"))
	assert.True(t, detector.Detect("when are romans being released?"))
	assert.True(t, detector.Detect("romans when?"))
	assert.True(t, detector.Detect("have the romans been released?"))
	assert.True(t, detector.Detect("romans out?"))
}
