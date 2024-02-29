package prepare

import (
	"github.com/lor08/goSkillbox/final/internal/models"
	"github.com/lor08/goSkillbox/final/internal/providers/voicecall"
	"log"
)

func GetVoiceCallData(c chan []models.VoiceCallData, countriesList map[string]models.Country) {
	defer close(c)
	voiceProvider := voicecall.VoiceProvider{
		Name: "Voice Calls",
	}
	voiceRes, err := voiceProvider.GetStatus(countriesList)
	if err != nil {
		log.Printf("can't get VoiceCalls status: %v", err)
		c <- nil
	}
	c <- voiceRes
}
