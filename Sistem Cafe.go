package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

type MenuItem struct {
    Name     string
    Price    float64
    Quantity int
}

type Category struct {
    Name  string
    Items []MenuItem
}

var orders []MenuItem

func main() {
    categories := []Category{
        {
            Name: "Minuman",
            Items: []MenuItem{
                {"Americano", 22000, 0},
                {"Latte", 25000, 0},
                {"Cappuccino", 23000, 0},
                {"Matcha Latte", 35000, 0},
                {"Leychee Tea", 26000, 0},
                {"Tropical Paradise", 32000, 0},
                {"Butterscotch Sea Salt", 31000, 0},
                {"Aren Latte", 29000, 0},
                {"Canggu Breeze", 25000, 0},
                {"Mineral Water", 10000, 0},
            },
        },
        {
            Name: "Makanan",
            Items: []MenuItem{
                {"Chicken Mentai", 35000, 0},
                {"Creamy Carbonara", 45000, 0},
                {"Aglio Olio", 44000, 0},
                {"Chicken Salad", 30000, 0},
                {"Javanese Fried Rice", 42000, 0},
                {"Potato Wedges", 26000, 0},
                {"Mix Platter", 30000, 0},
                {"Pain Au Tiramisu", 36000, 0},
                {"Almond Croissant", 33000, 0},
                {"Ayam Betutu", 45000, 0},
                {"Chicken Curry", 37000, 0},
            },
        },
    }

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("\n=== Menu Kafe ===")
        for i, cat := range categories {
            fmt.Printf("%d. %s\n", i+1, cat.Name)
        }
        fmt.Println("3. Selesai dan Cetak Nota")
        fmt.Print("Pilihan Anda: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        choice, _ := strconv.Atoi(input)
        if choice >= 1 && choice <= 2 {
            displayMenu(&categories[choice-1], reader)
        } else if choice == 3 {
            displayReceipt()
            break
        } else {
            fmt.Println("Pilihan tidak valid.")
        }
    }
}

func displayMenu(cat *Category, reader *bufio.Reader) {
    for {
        fmt.Printf("\n=== Daftar Menu %s ===\n", cat.Name)
        for i, item := range cat.Items {
            fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Name, item.Price)
        }
        fmt.Printf("%d. Sortir Harga ASC\n", len(cat.Items)+1)
        fmt.Printf("%d. Sortir Harga DESC\n", len(cat.Items)+2)
        fmt.Printf("%d. Cari Menu\n", len(cat.Items)+3)
        fmt.Printf("%d. Kembali\n", len(cat.Items)+4)

        fmt.Print("Pilih menu: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        choice, _ := strconv.Atoi(input)

        if choice >= 1 && choice <= len(cat.Items) {
            fmt.Printf("Jumlah '%s' yang ingin dipesan: ", cat.Items[choice-1].Name)
            qtyInput, _ := reader.ReadString('\n')
            qtyInput = strings.TrimSpace(qtyInput)
            qty, _ := strconv.Atoi(qtyInput)

            item := cat.Items[choice-1]
            item.Quantity = qty
            orders = append(orders, item)
            fmt.Printf("%d %s telah ditambahkan ke pesanan.\n", qty, item.Name)

        } else if choice == len(cat.Items)+1 {
            sort.Slice(cat.Items, func(i, j int) bool {
                return cat.Items[i].Price < cat.Items[j].Price
            })
            fmt.Println("Menu diurutkan dari harga terendah.")

        } else if choice == len(cat.Items)+2 {
            sort.Slice(cat.Items, func(i, j int) bool {
                return cat.Items[i].Price > cat.Items[j].Price
            })
            fmt.Println("Menu diurutkan dari harga tertinggi.")

        } else if choice == len(cat.Items)+3 {
            fmt.Print("Masukkan nama menu yang dicari: ")
            query, _ := reader.ReadString('\n')
            query = strings.TrimSpace(query)
            found := false
            for i, item := range cat.Items {
                if strings.Contains(strings.ToLower(item.Name), strings.ToLower(query)) {
                    fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Name, item.Price)
                    found = true
                }
            }
            if !found {
                fmt.Println("Menu tidak ditemukan.")
            }

        } else if choice == len(cat.Items)+4 {
            break
        } else {
            fmt.Println("Pilihan tidak valid.")
        }
    }
}

func displayReceipt() {
    fmt.Println("\n=======================================================================")
    fmt.Println("                            NOTA PEMBELIAN                              ")
    fmt.Println("=======================================================================")
    fmt.Printf("%-5s%-20s%-15s%-15s%-15s\n", "No", "Nama Item", "Harga", "Jumlah", "Total")

    total := 0.0
    for i, item := range orders {
        sub := item.Price * float64(item.Quantity)
        fmt.Printf("%-5d%-20sRp %-12.2f%-15dRp %-10.2f\n", i+1, item.Name, item.Price, item.Quantity, sub)
        total += sub
    }

    fmt.Println("-----------------------------------------------------------------------")
    fmt.Printf("%55s Rp %.2f\n", "Total Harga:", total)
    fmt.Println("=======================================================================")
    fmt.Println("              Hidangan terbaik untuk pelanggan istimewa                ")
    fmt.Println("                   Terima kasih telah berkunjung!                      ")
    fmt.Println("=======================================================================")
}
