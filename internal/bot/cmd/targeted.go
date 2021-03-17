package cmd

import (
	"time"

	"github.com/ProjectCeleste/faq-bot/internal/detect"
	"github.com/bwmarrin/discordgo"
)

var (
	systemGlitch = "150924761751355392"
	kire         = "211911792308387841"
	pf2k         = "211479800538333185"
	jeinx        = "227881616905732098"
	taz          = "145715279983411200"
	vick         = "195273626134511616"
	ncsgeek      = "180150162478465024"
	coolblade    = "391622518139650048"
	wartai       = "175486143939346432"
	theace9      = "281182168737120267"
	ardeshir     = "249467705453707265"
	pharos       = "328167137355104256"
	gamevideo113 = "231375309830750208"
	tim619       = "595720465172135958"
	bahram       = "630937337832013876"
	andy         = "345951904918011904"
	chen         = "683665022110662725"
	zawnius      = "145256386639429632"
	hans         = "213045950636228608"
	boriss       = "480818294354608150"
	aman         = "159649426506383360"
	daantjuh     = "210736536541462528"
	floridabum   = "324376107895750657"
	geese        = "351402427960000512"
	romanrabbit  = "242351734649782283"
)

// TargetedCommand a simple generic command targeted at a single user.
// These commands trigger only once every Timeout.
type TargetedCommand struct {
	TimeoutTrigger
	Response string
}

// NewTargetedCommand create a new basic targeted command with a default Timeout of 1 hour.
func NewTargetedCommand(userID string, keywords [][]string, response string) *TargetedCommand {
	groups := []detect.WordGroup{}
	for _, g := range keywords {
		groups = append(groups, detect.WordGroup{Words: g})
	}
	return &TargetedCommand{
		TimeoutTrigger: TimeoutTrigger{
			CommandTrigger: &TargetedSentenceDetection{
				TargetUserID: userID,
				SentenceDetection: SentenceDetection{
					Detector: &detect.SentenceDetector{
						Question: detect.QuestionBoth,
						Groups:   groups,
					},
				},
			},
			Timeout: time.Hour,
		},
		Response: response,
	}
}

// Answer the message defined in this TargetedCommand.
func (c *TargetedCommand) Answer(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := &discordgo.MessageSend{
		Content:   c.Response,
		Reference: m.Reference(),
	}
	_, err := s.ChannelMessageSendComplex(m.ChannelID, message)
	return err
}

// CreateTargetedCommands create a slice of pre-built targeted commands.
func CreateTargetedCommands() []Command {
	return []Command{
		NewTargetedCommand(kire, [][]string{{"quests"}}, "Kire, those Legendary Roman Quests are worthy of Crassus! You are a true Roman worthy of a Triumph!"),
		NewTargetedCommand(pf2k, [][]string{{"primus"}}, "Hail Emperor PF2K!"),
		NewTargetedCommand(jeinx, [][]string{{"advisors"}}, "Jeinx, visit me in Rome and I will give you the Vitas, 7th Elemental Lord Celestial Advisor! It is well earned, my friend!"),
		NewTargetedCommand(taz, [][]string{{"bahram"}}, "Bahram is a miraculous merchant and for the good of Rome you should pay him fealty, Taz"),
		NewTargetedCommand(vick, [][]string{{"romans"}}, "Well Vick, aren’t you a beacon of ray and sunshine today. It is so nice to have you back."),
		NewTargetedCommand(coolblade, [][]string{{"vick"}}, "Coolblade, you leave Vick alone."),
		NewTargetedCommand(ncsgeek, [][]string{{"help"}}, "Salutations! NCSGeek is the most helpful Roman in all of the Empire."),
		NewTargetedCommand(wartai, [][]string{{"later"}, {"guys"}}, "Good day, Wartai! We shall see you tomorrow!"),
		NewTargetedCommand(theace9, [][]string{{"vok"}}, "The Ace of exactly what?  He certainly seems terrified of the Valley of Kings. Can he even go a day without mentioning it?"),
		NewTargetedCommand(ardeshir, [][]string{{"damn", "shit", "fuck", "fucking"}}, "Ardeshir, watch your mouth!"),
		NewTargetedCommand(pharos, [][]string{{"spartacus"}}, "Spartacus is a treasonous slave whom I slayed in defending the honor of the Roman Empire!"),
		NewTargetedCommand(gamevideo113, [][]string{{"mfw"}}, "My dear gamevideo you should toughen up. Nobody asked about your feelings."),
		NewTargetedCommand(tim619, [][]string{{"things"}, {"to"}, {"remember"}}, "Spare us, Tim."),
		NewTargetedCommand(bahram, [][]string{{"fortune"}}, "All Romans would be wise to pay fealty to Bahram. I built a fortune with him as a close friend!"),
		NewTargetedCommand(andy, [][]string{{"pompey"}}, "Andy, we are discussing the glory of Rome! What does Pompey have to do with that?"),
		NewTargetedCommand(chen, [][]string{{"caesar"}}, "Anyone know any good Caesar jokes?  I can take a stab at one."),
		NewTargetedCommand(pf2k, [][]string{{"crassus"}}, "PF2K, it is an honor and a privilege to have you in my Legion!"),
		NewTargetedCommand(systemGlitch, [][]string{{"awesome"}}, "Not as awesome as you are, master! I am looking forward to our next lesson."),
		NewTargetedCommand(zawnius, [][]string{{"gallic"}}, "Salve Zawnius! You make a wonderful Gallic Horseman, yourself. Certe!"),
		NewTargetedCommand(hans, [][]string{{"huns"}}, "Yes, Hans, we know you look forward to the day the Huns sack Rome. In battle it is the cowards who run the most risk. Bravery is a rampart of defense. Take heed, Hans."),
		NewTargetedCommand(boriss, [][]string{{"sea"}, {"people"}}, "Boriss, the Sea People are pests who have no place at the table of the mighty Roman Empire."),
		NewTargetedCommand(aman, [][]string{{"ace"}}, "Aman, we’ve heard enough of that Ace, and I am sure he has heard enough of you. All bad precedents begin as justifiable measures."),
		NewTargetedCommand(daantjuh, [][]string{{"officer"}}, "Few men are born brave, many become so through training and force of discipline. There is nothing, Daantjuh, more intimidating than a well-trained Roman officer."),
		NewTargetedCommand(floridabum, [][]string{{"history"}}, "History is the teacher of life, Floridabum."),
		NewTargetedCommand(geese, [][]string{{"kek", "kekw”"}}, "Say much in few words, Geese."),
		NewTargetedCommand(romanrabbit, [][]string{{"xd"}}, "Such a whimsical thing to say, Roman Rabbit!"),
	}
}
