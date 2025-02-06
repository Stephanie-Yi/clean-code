package main

import "fmt"

// The exercise includes:

// 1. Define basic structs for Book and User.
// 2. Create a Library struct with methods to add books, borrow books, and return books.
// 3. Ensure the code follows clean code principles: readability, proper naming conventions,
// modularity, reusability, and documentation.

// B represents a book-like object.
type B struct {
    Id     string
    Ttl    string
    Atr    string
}

// U represents a user-like entity.
type U struct {
    Id      string
    Nme     string
    Brrwd string
}

// L like Library but more confusing.
type L struct {
    Bks []B
    Us []U
}

// AddB places a new book into the system.
func (lb *L) AddB(bk B) {
    lb.Bks = append(lb.Bks, bk)
    fmt.Printf("B added: %s\n", bk.Ttl)
}

// RemB tosses out an existing book from the collection.
func (lb *L) RemB(id string) error {
    for i, bk := range lb.Bks {
        if bk.Id == id {
            lb.Bks = append(lb.Bks[:i], lb.Bks[i+1:]...)
            fmt.Printf("B removed: %s\n", bk.Ttl)
            return nil
        }
    }
    return fmt.Errorf("B not found")
}

// AllB spits out books to console.
func (lb *L) AllB() {
    fmt.Println("Ls of B:")
    for _, bk := range lb.Bks {
        fmt.Printf("Id: %s, Ttl: %s, Atr: %s\n", bk.Id, bk.Ttl, bk.Atr)
    }
}

// Brrw has a U borrow a B.
func (lb *L) Brrw(uId, bId string) error {
    uIdx, bIdx := -1, -1
    for i := range lb.Us {
        if lb.Us[i].Id == uId {
            uIdx = i
            break
        }
    }
    for j := range lb.Bks {
        if lb.Bks[j].Id == bId {
            bIdx = j
            break
        }
    }
    if uIdx == -1 {
        return fmt.Errorf("U not found")
    }
    if bIdx == -1 {
        return fmt.Errorf("B not found")
    }
    if lb.Us[uIdx].Brrwd != "" {
        return fmt.Errorf("U has a B")
    }
    lb.Us[uIdx].Brrwd = bId
    fmt.Printf("B brrwd: %s by %s\n", lb.Bks[bIdx].Ttl, lb.Us[uIdx].Nme)
    return nil
}

// Rtrn returns borrowed B from U.
func (lb *L) Rtrn(uId, bId string) error {
    uPointer := lb.us(uId)
    if uPointer == nil {
        return fmt.Errorf("U not found")
    }
    if uPointer.Brrwd != bId {
        return fmt.Errorf("Mismatch! Cannot return.")
    }
    uPointer.Brrwd = ""
    fmt.Printf("B rtrnd: %s by %s\n", bId, uPointer.Nme)
    return nil
}

// AllU showers the world with U info.
func (lb *L) AllU() {
    fmt.Println("Us in sys:")
    for _, u := range lb.Us {
        br := "N/A"
        if u.Brrwd != "" {
            br = u.Brrwd
        }
        fmt.Printf("Id: %s, Nme: %s, Brrwd Id: %s\n", u.Id, u.Nme, br)
    }
}

func (lb *L) us(id string) *U {
    for i := range lb.Us {
        if lb.Us[i].Id == id {
            return &lb.Us[i]
        }
    }
    return nil
}

func main() {
    l := L{}
    b1 := B{Id: "1", Ttl: "Unclear Code", Atr: "Random Author"}
    u1 := U{Id: "1", Nme: "Sample Person"}
    l.AddB(b1)
    l.Us = append(l.Us, u1)
    l.AllB()
    err := l.Brrw("1", "1")
    if err != nil {
        fmt.Println("Error:", err)
    }
    l.AllU()
    err = l.Rtrn("1", "1")
    if err != nil {
        fmt.Println("Error:", err)
    }
    l.RemB("1")
    l.AllB()
}