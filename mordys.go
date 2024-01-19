package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

// tea "github.com/charmbracelet/bubbletea"

// type model struct {
// 	cursor   int
// 	choices  []string
// 	selected map[int]struct{}
// }

// func initialModel() model {
// 	return model{
// 		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

// 		// A map which indicates which choices are selected. We're using
// 		// the map like a mathematical set. The keys refer to the indexes
// 		// of the `choices` slice, above.
// 		selected: make(map[int]struct{}),
// 	}
// }

// func (m model) Init() tea.Cmd {
// 	return tea.SetWindowTitle("Grocery List")
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q":
// 			return m, tea.Quit
// 		case "up", "k":
// 			if m.cursor > 0 {
// 				m.cursor--
// 			}
// 		case "down", "j":
// 			if m.cursor < len(m.choices)-1 {
// 				m.cursor++
// 			}
// 		case "enter", " ":
// 			_, ok := m.selected[m.cursor]
// 			if ok {
// 				delete(m.selected, m.cursor)
// 			} else {
// 				m.selected[m.cursor] = struct{}{}
// 			}
// 		}
// 	}

// 	return m, nil
// }

// func (m model) View() string {
// 	s := "What should we buy at the market?\n\n"

// 	for i, choice := range m.choices {
// 		cursor := " "
// 		if m.cursor == i {
// 			cursor = ">"
// 		}

// 		checked := " "
// 		if _, ok := m.selected[i]; ok {
// 			checked = "x"
// 		}

// 		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
// 	}

// 	s += "\nPress q to quit.\n"

// 	return s
// }

// func main() {
// 	p := tea.NewProgram(initialModel())
// 	if _, err := p.Run(); err != nil {
// 		fmt.Printf("Alas, there's been an error: %v", err)
// 		os.Exit(1)
// 	}
// }

// Larian Studios Pak Header
type LSPKHeader struct {
	Version uint32
	FileListOffset uint64
	FileListSize uint32
	Flags byte
	Priority byte
	Md5 [16]byte
	NumParts uint16
}

// Larian Studios Pak Sub-File Entry
type LSPKFileEntry struct {
	Name [256]byte
	OffsetInFile1 uint32
	OffsetInFile2 uint16
	ArchivePart byte
	Flags byte
	SizeOnDisk uint32
	UncompressedSize uint32
}



func ReadSignature(reader *bytes.Reader) ([]byte){
	buf := make([]byte, 4)
	io.ReadAtLeast(reader, buf, 4)
	return buf
}

func ReadHeader(reader *bytes.Reader) (*LSPKHeader, error) {
	var header LSPKHeader
	reader.Seek(4, 0)
	err := binary.Read(reader, binary.LittleEndian, &header)
    if err != nil {
        log.Fatal("read error:", err)
    }
	return &header, nil
}

func main() {
	data, err := os.ReadFile("./test_files/5eSpells.pak")
	if err != nil {
        panic(err)
    }

    reader := bytes.NewReader(data)
	
	sig := ReadSignature(reader)
	fmt.Printf("%s\n", sig)
	
    header, err := ReadHeader(reader)
	if err != nil {
        panic(err)
    }
	fmt.Printf("header: %v\n", header)
	
}