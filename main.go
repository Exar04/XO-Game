package main

import (
	"fmt"
	"os"
//	"time"

	tea "github.com/charmbracelet/bubbletea"
//	"golang.org/x/tools/go/analysis/passes/ifaceassert"
)

type model struct {
        cursorX int
        cursorY int
        choices [][]int
        selected [][]int
        playerTurn int
}

func initialModel() model {
    n := 3 // replace with your desired size
    selected := make([][]int, n)
    for i := range selected {
        selected[i] = make([]int, n)
    }

    choices := make([][]int, n)
    for i := range choices {
        choices[i] = make([]int, n)
    }

	return model{
        selected: selected,
        choices: choices,
        playerTurn: 1,
	}
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursorX > 0 {
                m.cursorX--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursorX < 2 {
                m.cursorX++
            }
        
        case "left", "h":
            if m.cursorY > 0  {
                m.cursorY--
            }

        case "right", "l":
            if m.cursorY < 2 {
                m.cursorY++
            }

        // The "enter" key and the spacebar (a literal space) toggle
        // the selected state for the item that the cursor is pointing at.
        case "enter", " ":
            // cursorPos := m.cursorX+m.cursorY
            PositionVacant := m.selected[m.cursorX][m.cursorY]

            player := m.playerTurn
             if player == 1 && PositionVacant == 0{
                 m.selected[m.cursorX][m.cursorY] = 1
                 m.playerTurn = 2

             } else if player == 2 && PositionVacant == 0 {
                 m.selected[m.cursorX][m.cursorY] = 3
                 m.playerTurn = 1
             }
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
    s := "start playing \n"
    for i:= 0; i < 3; i++ {
        for j:= 0; j < 3; j++{

            // Is the cursor pointing at this choice?

            // Is this choice selected?
            checked := " " // not selected
            ok := m.selected[i][j];
            //        player := m.playerTurn
            if  ok==1 {
                checked = "x" // selected!
            } else if ok == 3 {
                checked = "o"
            } else {
                checked = " "
            }

            cursor := " " // no cursor

            // Render the row
            if m.cursorX == i && m.cursorY == j {
                cursor = ">" // cursor!
            }

            s += fmt.Sprintf("%s [%s] ", cursor, checked)
        }
        s += fmt.Sprintf("\n")
    }

    // The footer
    s += fmt.Sprintf("\n Player %d plays\n", m.playerTurn)
    if checkWhoWon(m.selected)!="DRAW"{
        s = "Game Over\n"
        s += checkWhoWon(m.selected)

        deleteArrayValues(m.selected)
        deleteArrayValues(m.choices)
//        time.Sleep(1 * time.Second)
    }
    s += "\nPress q to quit.\n"
return s
}

func checkWhoWon(arr [][]int)string{
    var P1won int
    var P2won int
    for i := range arr {
        for j := range arr{
            if arr[i][j] == 1 {
                P1won += arr[i][j]
            }

            if arr[i][j] == 3 {
                P2won += arr[i][j]
            }

            if P1won == 3 {
                return "Player 1 won"
            }
            if P2won == 9 {
                return "Player 2 won"
            }
        }
        P1won = 0
        P2won = 0
    }
    for i := range arr {
        for j := range arr{
            if arr[j][i] == 1 {
                P1won += arr[j][i]
            }
            if arr[j][i] == 3 {
                P2won += arr[j][i]
            }
            
            if P1won == 3 {
                return "Player 1 won"
            }
            if P2won == 9 {
                return "Player 2 won"
            }
        }
        P1won = 0
        P2won = 0
    }
    return "DRAW" 
}

func deleteArrayValues(arr [][]int){
     for i := range arr {
        for j := range arr{
            arr[i][j] = 0
        }
    }
}
func ArrayIsFull(arr [][]int)bool{
    for i := range arr {
        for j := range arr{
            if arr[i][j] == 0 {
                return false
            }
        }
    }
    return true
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
