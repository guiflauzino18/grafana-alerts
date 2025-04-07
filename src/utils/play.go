package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"github.com/go-audio/wav"

	"github.com/hajimehoshi/oto/v2"
)

func PlayAudio(arquivo string) error {
	f, erro := os.Open(arquivo)
	if erro != nil {
		return fmt.Errorf("Erro ao abrir arquivo de audio: \n%w", erro)
	}

	defer f.Close()

	decoder := wav.NewDecoder(f)
	if !decoder.IsValidFile() {
		return fmt.Errorf("Arquivo wav invalido")
	}

	buf, erro := decoder.FullPCMBuffer()
	if erro != nil {
		return fmt.Errorf("Erro ao ler Buffer:\n%w", erro)
	}

	var data = new(bytes.Buffer)
	for _, sample := range buf.Data {
		if erro := binary.Write(data, binary.LittleEndian, int16(sample)); erro != nil {
			return fmt.Errorf("Erro ao converter sample:\n%w", erro)
		}
	}

	contexto, ready, erro := oto.NewContext(int(decoder.SampleRate), int(decoder.NumChans), 2)
	if erro != nil {
		return fmt.Errorf("Erro ao criar contexto oto:\n%w", erro)
	}

	<-ready

	player := contexto.NewPlayer(data)
	defer player.Close()

	player.Play()

	for player.IsPlaying() {
		time.Sleep(10 * time.Millisecond)
	}
	return nil

}
