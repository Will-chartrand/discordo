package config

import "github.com/rivo/tview"

type (
	Theme struct {
		Border        bool   `toml:"border"`
		BorderColor   string `toml:"border_color"`
		BorderPadding [4]int `toml:"border_padding"`

		TitleColor      string `toml:"title_color"`
		BackgroundColor string `toml:"background_color"`

		GuildsTree   GuildsTreeTheme   `toml:"guilds_tree"`
		MessagesText MessagesTextTheme `toml:"messages_text"`
	}

	GuildsTreeTheme struct {
		AutoExpandFolders   bool   `toml:"auto_expand_folders"`
		ChannelColor        string `toml:"channel_color"`
		Graphics            bool   `toml:"graphics"`
		GuildColor          string `toml:"guild_color"`
		PrivateChannelColor string `toml:"private_channel_color"`
	}

	MessagesTextTheme struct {
		ReplyIndicator string `toml:"reply_indicator"`

		AuthorColor     string `toml:"author_color"`
		OtherColor      string `toml:"other_color"`
		ContentColor    string `toml:"content_color"`
		EmojiColor      string `toml:"emoji_color"`
		LinkColor       string `toml:"link_color"`
		AttachmentColor string `toml:"attachment_color"`
	}
)

func defaultTheme() Theme {
	return Theme{
		Border:        true,
		BorderColor:   "default",
		BorderPadding: [...]int{0, 0, 1, 1},

		BackgroundColor: "default",
		TitleColor:      "default",

		GuildsTree: GuildsTreeTheme{
			AutoExpandFolders:   true,
			ChannelColor:        tview.Styles.PrimaryTextColor.String(),
			Graphics:            true,
			GuildColor:          tview.Styles.PrimaryTextColor.String(),
			PrivateChannelColor: tview.Styles.PrimaryTextColor.String(),
		},
		MessagesText: MessagesTextTheme{
			ReplyIndicator: string(tview.BoxDrawingsLightArcDownAndRight) + " ",

			AuthorColor:     "green",
			OtherColor:      "aqua",
			ContentColor:    tview.Styles.PrimaryTextColor.String(),
			EmojiColor:      "green",
			LinkColor:       "blue",
			AttachmentColor: "yellow",
		},
	}
}
