package cmd

import "github.com/ProjectCeleste/faq-bot/internal/detect"

// NewGameNotAvailable create an assistance command for install without steam
func NewGameNotAvailable() *GlobalCommand {
	return &GlobalCommand{
		TimeoutTrigger: TimeoutTrigger{
			CommandTrigger: &SentenceDetection{
				Detector: &detect.SentenceDetector{
					Question: detect.QuestionBoth,
					Groups: []detect.WordGroup{
						{Words: []string{"can't", "cannot"}},
						{Words: []string{"download", "install"}},
						{Words: []string{"steam"}},
					},
					Variants: []detect.SentenceDetector{
						{
							Question: detect.QuestionBoth,
							Groups: []detect.WordGroup{
								{Words: []string{"game"}},
								{Words: []string{"not"}},
								{Words: []string{"available"}},
								{Words: []string{"region"}},
							},
						},
						{
							Question: detect.QuestionBoth,
							Groups: []detect.WordGroup{
								{Words: []string{"not"}},
								{Words: []string{"available"}},
								{Words: []string{"steam"}},
							},
						},
						{
							Question: detect.QuestionBoth,
							Groups: []detect.WordGroup{
								{Words: []string{"steam"}},
								{Words: []string{"not"}},
								{Words: []string{"available"}},
							},
						},
					},
				},
			},
		},
		Response: "Greetings! You can download the game using the non-Steam option here: <https://www.projectceleste.com/install#install_the_original_game_without_using_steam>",
	}
}

// NewUpdateLauncher create an assistance command for launcher updates.
func NewUpdateLauncher() *GlobalCommand {
	return &GlobalCommand{
		TimeoutTrigger: TimeoutTrigger{
			CommandTrigger: &SentenceDetection{
				Detector: &detect.SentenceDetector{
					Question: detect.QuestionOnly,
					Groups: []detect.WordGroup{
						{Words: []string{"where", "how"}},
						{Words: []string{"download", "update", "get"}},
						{Words: []string{"launcher"}},
					},
					Variants: []detect.SentenceDetector{
						{
							Question: detect.QuestionOnly,
							Groups: []detect.WordGroup{
								{Words: []string{"where", "how"}},
								{Words: []string{"download", "update", "get"}},
								{Words: []string{"latest", "last"}},
								{Words: []string{"launcher"}},
								{Words: []string{"version"}},
							},
						},
						{
							Question: detect.QuestionOnly,
							Groups: []detect.WordGroup{
								{Words: []string{"where", "how"}},
								{Words: []string{"download", "update", "get"}},
								{Words: []string{"latest", "last"}},
								{Words: []string{"version"}},
								{Words: []string{"launcher"}},
							},
						},
					},
				},
			},
		},
		Response: "Greetings and welcome! Click on the following link then \"Assets -> CelesteLauncher.zip\" to download the newest version of the Launcher: <https://github.com/ProjectCeleste/Celeste_Launcher/releases/latest>",
	}
}

// NewInstallLinux create an assistance command for installing the game on linux.
func NewInstallLinux() *GlobalCommand {
	return &GlobalCommand{
		TimeoutTrigger: TimeoutTrigger{
			CommandTrigger: &SentenceDetection{
				Detector: &detect.SentenceDetector{
					Question: detect.QuestionOnly,
					Groups: []detect.WordGroup{
						{Words: []string{"how", "can"}},
						{Words: []string{"install", "run"}},
						{Words: []string{"linux"}},
					},
					Variants: []detect.SentenceDetector{
						{
							Question: detect.QuestionOnly,
							Groups: []detect.WordGroup{
								{Words: []string{"game"}},
								{Words: []string{"run", "work"}},
								{Words: []string{"linux"}},
							},
						},
					},
				},
			},
		},
		Response: "Greetings and welcome! The game can be installed on linux using our Lutris installer: <https://www.projectceleste.com/install/#install_on_linux>",
	}
}

// NewGameDoesntStart create an assistance command for troubleshooting.
func NewGameDoesntStart() *GlobalCommand {
	return &GlobalCommand{
		TimeoutTrigger: TimeoutTrigger{
			CommandTrigger: &SentenceDetection{
				Detector: &detect.SentenceDetector{
					Question: detect.QuestionBoth,
					Groups: []detect.WordGroup{
						{Words: []string{"game"}},
						{Words: []string{"don't", "doesn't"}},
						{Words: []string{"start", "run"}},
					},
					Variants: []detect.SentenceDetector{
						{
							Question: detect.QuestionBoth,
							Groups: []detect.WordGroup{
								{Words: []string{"can't", "cannot"}},
								{Words: []string{"start", "run"}},
								{Words: []string{"game"}},
							},
						},
					},
				},
			},
		},
		Response: "Well met! You may want to try these troubleshooting steps: <https://www.projectceleste.com/install/#troubleshooting>\n\nIf that doesn't help, ask in our <#322275547264581632> section and one of our tech support will help you soon.",
	}
}

// NewChatEmpty create an assistance command for empty chat.
func NewChatEmpty() *GlobalCommand {
	return &GlobalCommand{
		TimeoutTrigger: TimeoutTrigger{
			CommandTrigger: &SentenceDetection{
				Detector: &detect.SentenceDetector{
					Question: detect.QuestionBoth,
					Groups: []detect.WordGroup{
						{Words: []string{"cannot", "can't"}},
						{Words: []string{"see"}},
						{Words: []string{"anything"}},
						{Words: []string{"chat"}},
					},
					Variants: []detect.SentenceDetector{
						{
							Question: detect.QuestionBoth,
							Groups: []detect.WordGroup{
								{Words: []string{"chat"}},
								{Words: []string{"is"}},
								{Words: []string{"empty"}},
							},
						},
						{
							Question: detect.QuestionBoth,
							Groups: []detect.WordGroup{
								{Words: []string{"nothing"}},
								{Words: []string{"in"}},
								{Words: []string{"chat"}},
							},
						},
					},
				},
			},
		},
		Response: "Make sure you have the Chat Channel enabled. To do so, you can right-click on any of the chat tabs on the bottom-left chatbox and click on \"Options\".\nhttps://cdn.discordapp.com/attachments/488086268773924883/821340565861040128/unknown.png",
	}
}

// NewChangeUsername create an assistance command for changing username.
func NewChangeUsername() *GlobalCommand {
	return &GlobalCommand{
		TimeoutTrigger: TimeoutTrigger{
			CommandTrigger: &SentenceDetection{
				Detector: &detect.SentenceDetector{
					Question: detect.QuestionOnly,
					Groups: []detect.WordGroup{
						{Words: []string{"how", "can"}},
						{Words: []string{"change"}},
						{Words: []string{"username", "name"}},
					},
				},
			},
		},
		Response: "Greetings! You may send a message to one of the Moderators with your current name in-game, and your desired name. If the username isn't already taken, you will get your account renamed. Remember, one rename per account!",
	}
}

// NewQuestCrashes create an assistance command for when quests crash and users need to abandon and retake them.
func NewQuestCrashes() *GlobalCommand {
	return &GlobalCommand{
		TimeoutTrigger: TimeoutTrigger{
			CommandTrigger: &SentenceDetection{
				Detector: &detect.SentenceDetector{
					Question: detect.QuestionBoth,
					Groups: []detect.WordGroup{
						{Words: []string{"quest"}},
						{Words: []string{"crash", "exit"}},
					},
					Variants: []detect.SentenceDetector{
						{
							Question: detect.QuestionBoth,
							Groups: []detect.WordGroup{
								{Words: []string{"mission"}},
								{Words: []string{"objective", "crashes", "crashes without errors", "exits with no error", "no error", "hard crash", "bugged", "exit"}},
							},
						},
					},
				},
			},
		},
		Response: "Greetings! Try abandoning your quest from your Quest Log and taking it again from the Quest Giver! In most cases, this solves the crashing issue for quests.",
	}
}
