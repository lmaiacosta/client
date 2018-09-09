// Copyright 2018 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"encoding/json"
	"time"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/chat1"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	"golang.org/x/net/context"
)

type CmdChatAPIListen struct {
	libkb.Contextified
}

func (c *CmdChatAPIListen) ParseArgv(ctx *cli.Context) error {
	return nil
}

func newCmdChatAPIListen(cl *libcmdline.CommandLine, g *libkb.GlobalContext) cli.Command {
	return cli.Command{
		Name: "api-listen",
		// No "Usage" field makes it hidden in command list.
		Description: "Listen and print incoming chat actions in JSON format",
		Action: func(c *cli.Context) {
			cl.ChooseCommand(&CmdChatAPIListen{
				Contextified: libkb.NewContextified(g),
			}, "api-listen", c)
			cl.SetNoStandalone()
		},
		Flags: []cli.Flag{},
	}
}

func NewCmdChatAPIListenRunner(g *libkb.GlobalContext) *CmdChatAPIListen {
	return &CmdChatAPIListen{
		Contextified: libkb.NewContextified(g),
	}
}

func (c *CmdChatAPIListen) Run() error {
	display := &chatNotificationDisplay{
		Contextified: libkb.NewContextified(c.G()),
	}
	protocols := []rpc.Protocol{
		chat1.NotifyChatProtocol(display),
	}
	channels := keybase1.NotificationChannels{
		Chat: true,
	}

	if err := RegisterProtocolsWithContext(protocols, c.G()); err != nil {
		return err
	}
	cli, err := GetNotifyCtlClient(c.G())
	if err != nil {
		return err
	}
	if err := cli.SetNotifications(context.TODO(), channels); err != nil {
		return err
	}

	jsonStr, _ := json.Marshal(struct {
		What string `json:"what"`
	}{
		What: "listening",
	})
	display.printf("%s\n", jsonStr)
	for {
		time.Sleep(time.Second)
	}
}

func (c *CmdChatAPIListen) GetUsage() libkb.Usage {
	return libkb.Usage{}
}

type chatNotificationDisplay struct {
	libkb.Contextified
}

func (d *chatNotificationDisplay) printf(fmt string, args ...interface{}) error {
	_, err := d.G().UI.GetTerminalUI().Printf(fmt, args...)
	return err
}

type incomingMessage struct {
	ConvName string `json:"conv_name"`
	Channel  string `json:"channel,omitempty"`
	Type     string `json:"type"`
	Message  string `json:"message,omitempty"`
	Sender   string `json:"sender,omitempty"`
}

func (d *chatNotificationDisplay) NewChatActivity(ctx context.Context, arg chat1.NewChatActivityArg) error {
	activity := arg.Activity
	typ, err := activity.ActivityType()
	if err == nil {
		switch typ {
		case chat1.ChatActivityType_INCOMING_MESSAGE:
			inMsg := activity.IncomingMessage()
			if inMsg.Message.IsValid() {
				mv := inMsg.Message.Valid()
				msgJSON := incomingMessage{
					ConvName: inMsg.Conv.Name,
					Channel:  inMsg.Conv.Channel,
					Sender:   mv.SenderUsername,
				}
				bodyType, err := mv.MessageBody.MessageType()
				if err == nil {
					msgJSON.Type = bodyType.String()
					switch bodyType {
					case chat1.MessageType_TEXT:
						msgJSON.Message = mv.MessageBody.Text().Body
					}
				}
				jsonStr, _ := json.Marshal(msgJSON)
				d.printf("%s\n", string(jsonStr))
			}
		}
	}
	return nil
}

func (d *chatNotificationDisplay) ChatIdentifyUpdate(context.Context, keybase1.CanonicalTLFNameAndIDWithBreaks) error {
	return nil
}
func (d *chatNotificationDisplay) ChatTLFFinalize(context.Context, chat1.ChatTLFFinalizeArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatTLFResolve(context.Context, chat1.ChatTLFResolveArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatInboxStale(context.Context, keybase1.UID) error { return nil }
func (d *chatNotificationDisplay) ChatThreadsStale(context.Context, chat1.ChatThreadsStaleArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatTypingUpdate(context.Context, []chat1.ConvTypingUpdate) error {
	return nil
}
func (d *chatNotificationDisplay) ChatJoinedConversation(context.Context, chat1.ChatJoinedConversationArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatLeftConversation(context.Context, chat1.ChatLeftConversationArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatResetConversation(context.Context, chat1.ChatResetConversationArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatInboxSyncStarted(context.Context, keybase1.UID) error {
	return nil
}
func (d *chatNotificationDisplay) ChatInboxSynced(context.Context, chat1.ChatInboxSyncedArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatSetConvRetention(context.Context, chat1.ChatSetConvRetentionArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatSetTeamRetention(context.Context, chat1.ChatSetTeamRetentionArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatSetConvSettings(context.Context, chat1.ChatSetConvSettingsArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatSubteamRename(context.Context, chat1.ChatSubteamRenameArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatKBFSToImpteamUpgrade(context.Context, chat1.ChatKBFSToImpteamUpgradeArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatAttachmentUploadStart(context.Context, chat1.ChatAttachmentUploadStartArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatAttachmentUploadProgress(context.Context, chat1.ChatAttachmentUploadProgressArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatPaymentInfo(context.Context, chat1.ChatPaymentInfoArg) error {
	return nil
}
func (d *chatNotificationDisplay) ChatRequestInfo(context.Context, chat1.ChatRequestInfoArg) error {
	return nil
}