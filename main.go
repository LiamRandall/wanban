//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world hello --out gen ./wit
package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"go.wasmcloud.dev/component/net/wasihttp"
)

// Since we don't run this program like a CLI, the `main` function is empty. Instead,
// we call the `handleRequest` function when an HTTP request is received.
func main() {}

// Data Structures
type Card struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type List struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Cards []Card `json:"cards"`
}

type Board struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Lists []List `json:"lists"`
}

// In-memory store
var boards []Board

// Utility function to write JSON responses
func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// Handle functions
func handleBoards(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// List all boards
		writeJSON(w, boards)
	case "POST":
		// Create a new board
		var b Board
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		b.ID = uuid.New().String()
		boards = append(boards, b)
		writeJSON(w, b)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleBoard(w http.ResponseWriter, r *http.Request, id string) {
	switch r.Method {
	case "GET":
		// Get a single board by ID
		for _, b := range boards {
			if b.ID == id {
				writeJSON(w, b)
				return
			}
		}
		http.NotFound(w, r)
	case "DELETE":
		// Delete a board
		for i, b := range boards {
			if b.ID == id {
				boards = append(boards[:i], boards[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.NotFound(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleLists(w http.ResponseWriter, r *http.Request, boardID string) {
	switch r.Method {
	case "GET":
		// Get lists for a board
		for _, b := range boards {
			if b.ID == boardID {
				writeJSON(w, b.Lists)
				return
			}
		}
		http.NotFound(w, r)
	case "POST":
		// Create a new list for a board
		var l List
		if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		l.ID = uuid.New().String()
		for i, b := range boards {
			if b.ID == boardID {
				boards[i].Lists = append(boards[i].Lists, l)
				writeJSON(w, l)
				return
			}
		}
		http.NotFound(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleList(w http.ResponseWriter, r *http.Request, boardID, listID string) {
	switch r.Method {
	case "GET":
		// Get a single list
		for _, b := range boards {
			if b.ID == boardID {
				for _, l := range b.Lists {
					if l.ID == listID {
						writeJSON(w, l)
						return
					}
				}
			}
		}
		http.NotFound(w, r)
	case "DELETE":
		// Delete a list from a board
		for bi, b := range boards {
			if b.ID == boardID {
				for li, l := range b.Lists {
					if l.ID == listID {
						boards[bi].Lists = append(boards[bi].Lists[:li], boards[bi].Lists[li+1:]...)
						w.WriteHeader(http.StatusNoContent)
						return
					}
				}
			}
		}
		http.NotFound(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleCards(w http.ResponseWriter, r *http.Request, boardID, listID string) {
	switch r.Method {
	case "GET":
		// Get cards for a list
		for _, b := range boards {
			if b.ID == boardID {
				for _, l := range b.Lists {
					if l.ID == listID {
						writeJSON(w, l.Cards)
						return
					}
				}
			}
		}
		http.NotFound(w, r)
	case "POST":
		// Create a new card
		var c Card
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		c.ID = uuid.New().String()

		for bi, b := range boards {
			if b.ID == boardID {
				for li, l := range b.Lists {
					if l.ID == listID {
						boards[bi].Lists[li].Cards = append(boards[bi].Lists[li].Cards, c)
						writeJSON(w, c)
						return
					}
				}
			}
		}
		http.NotFound(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleCard(w http.ResponseWriter, r *http.Request, boardID, listID, cardID string) {
	switch r.Method {
	case "GET":
		// Get a single card
		for _, b := range boards {
			if b.ID == boardID {
				for _, l := range b.Lists {
					if l.ID == listID {
						for _, c := range l.Cards {
							if c.ID == cardID {
								writeJSON(w, c)
								return
							}
						}
					}
				}
			}
		}
		http.NotFound(w, r)
	case "DELETE":
		// Delete a card
		for bi, b := range boards {
			if b.ID == boardID {
				for li, l := range b.Lists {
					if l.ID == listID {
						for ci, c := range l.Cards {
							if c.ID == cardID {
								boards[bi].Lists[li].Cards = append(boards[bi].Lists[li].Cards[:ci], boards[bi].Lists[li].Cards[ci+1:]...)
								w.WriteHeader(http.StatusNoContent)
								return
							}
						}
					}
				}
			}
		}
		http.NotFound(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//func init() {
// Register the handleRequest function as the handler for all incoming requests.
//	wasihttp.HandleFunc(handleRequest)
//}

//func init() {
// Register the handleRequest function as the handler for all incoming requests.
//	wasihttp.HandleFunc(handleRequest)
//}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello from Go!\n")
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	switch {
	case r.URL.Path == "/boards" && len(parts) == 1:
		handleBoards(w, r)
	case len(parts) == 2 && parts[0] == "boards":
		boardID := parts[1]
		handleBoard(w, r, boardID)
	case len(parts) == 3 && parts[0] == "boards" && parts[2] == "lists":
		boardID := parts[1]
		handleLists(w, r, boardID)
	case len(parts) == 4 && parts[0] == "boards" && parts[2] == "lists":
		boardID := parts[1]
		listID := parts[3]
		handleList(w, r, boardID, listID)
	case len(parts) == 5 && parts[0] == "boards" && parts[2] == "lists" && parts[4] == "cards":
		boardID := parts[1]
		listID := parts[3]
		handleCards(w, r, boardID, listID)
	case len(parts) == 6 && parts[0] == "boards" && parts[2] == "lists" && parts[4] == "cards":
		boardID := parts[1]
		listID := parts[3]
		cardID := parts[5]
		handleCard(w, r, boardID, listID, cardID)
	default:
		http.NotFound(w, r)
	}
}

//		log.Println("Server started on :8080")
//		log.Fatal(http.ListenAndServe(":8080", nil))
//}

// Basic router
func init() {
	wasihttp.HandleFunc(handleRequest)

}
