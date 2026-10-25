package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHasTriggerUsesUnicodeWordBoundaries(t *testing.T) {
	if !hasTrigger("Ну что, осуждаю!", []string{"осуждаю"}) {
		t.Fatal("expected trigger to match")
	}
	if hasTrigger("осуждающе", []string{"осуждаю"}) {
		t.Fatal("expected embedded word not to match")
	}
}

func TestLoadSettingsFromExternalConfig(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.cfg")
	contents := `[telegram]
bot_token = token
relay_forwarded_messages = true
relay_chat_id = -100123

[deepseek]
api_key = key
base_url = https://api.example.test
model = test-model
temperature = 0.4
max_tokens = 80
request_timeout = 12

[bot]
trigger_words = осуждаю, ну да
respond_to_direct_messages = true
bot_reference_words = бот
context_messages_after_reply = 2
context_reset_seconds = 45
cooldown_seconds = 3
max_message_chars = 100
system_prompt = Test prompt
`
	if err := os.WriteFile(path, []byte(contents), 0600); err != nil {
		t.Fatal(err)
	}
	settings, err := loadSettings(path)
	if err != nil {
		t.Fatal(err)
	}
	if settings.RelayChatID != -100123 || !settings.RelayForwardedMessages {
		t.Fatalf("unexpected relay settings: %+v", settings)
	}
	if settings.DeepSeekModel != "test-model" || settings.MaxTokens != 80 {
		t.Fatalf("unexpected DeepSeek settings: %+v", settings)
	}
	if len(settings.TriggerWords) != 2 || settings.TriggerWords[1] != "ну да" {
		t.Fatalf("unexpected trigger words: %#v", settings.TriggerWords)
	}
}

func TestNormalizeForwardText(t *testing.T) {
	got := normalizeForwardText("  Привет\n\tМИР  ")
	if got != "привет мир" {
		t.Fatalf("unexpected normalized text: %q", got)
	}
}
