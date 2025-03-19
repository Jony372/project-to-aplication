package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InsertRows(ctx context.Context, tx pgx.Tx, item Item) error {
    // Insert four rows into the "accounts" table.
    log.Println("Creating new row...")
    if _, err := tx.Exec(ctx,
        "INSERT INTO item (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", 
        item.Ticker,
        item.TargetFrom,
        item.TargetTo,
        item.Company,
        item.Action,
        item.Brokerage,
        item.RatingFrom,
        item.RatingTo,
        item.Time);
    err != nil {
        return err
    }
    return nil
}

func GrowthPorcentage(ctx context.Context, tx pgx.Tx, id string) (float64, error) {
    var porcent float64
    err := tx.QueryRow(ctx,
        `SELECT
            (target_to_num * 100 / target_from_num)-100 AS porcent
        FROM (
            SELECT
                id,
                CAST(REPLACE(REPLACE(target_to, '$', ''), ',', '') AS NUMERIC) AS target_to_num,
                CAST(REPLACE(REPLACE(target_from, '$', ''), ',', '') AS NUMERIC) AS target_from_num
            FROM item
            WHERE id = $1
        ) AS subquery
        WHERE target_to_num > target_from_num
        ORDER BY porcent DESC
        LIMIT 1;`,id).Scan(&porcent)
    if err != nil {
        return 0, err
    }
    return porcent, nil
}

func GetGrowthItems(ctx context.Context, tx pgx.Tx) ([]ChartItem, error) {
    var items []ChartItem
    rows, err := tx.Query(ctx,
        `SELECT
            id,
            ticker,
            (target_to_num * 100 / target_from_num)-100 AS porcent
        FROM (
            SELECT
                id,
                ticker,
                CAST(REPLACE(REPLACE(target_to, '$', ''), ',', '') AS NUMERIC) AS target_to_num,
                CAST(REPLACE(REPLACE(target_from, '$', ''), ',', '') AS NUMERIC) AS target_from_num
            FROM item
        ) AS subquery
        WHERE target_to_num > target_from_num
        ORDER BY porcent DESC
        LIMIT 10;`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var item ChartItem
        // Escanear cada fila en la estructura `item`
        if err := rows.Scan(&item.Id, &item.Ticker, &item.Porcent); err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}

func GetToBuyItems(ctx context.Context, tx pgx.Tx) ([]ChartItem, error) {
    var items []ChartItem
    rows, err := tx.Query(ctx,
        `SELECT
            id,
            ticker,
            (target_to_num * 100 / target_from_num)-100 AS porcent
        FROM (
            SELECT
                id,
                ticker,
                rating_to,
                CAST(REPLACE(REPLACE(target_to, '$', ''), ',', '') AS NUMERIC) AS target_to_num,
                CAST(REPLACE(REPLACE(target_from, '$', ''), ',', '') AS NUMERIC) AS target_from_num
            FROM item
        ) AS subquery
        WHERE rating_to = 'Buy'
        ORDER BY porcent ASC
        LIMIT 10;`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var item ChartItem
        // Escanear cada fila en la estructura `item`
        if err := rows.Scan(&item.Id, &item.Ticker, &item.Porcent); err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}

func GetDecreaseItems(ctx context.Context, tx pgx.Tx) ([]ChartItem, error) {
    var items []ChartItem
    rows, err := tx.Query(ctx,
        `SELECT
            id,
            ticker,
            100-(target_to_num * 100 / target_from_num) AS porcent
        FROM (
            SELECT
                id, ticker,
                CAST(REPLACE(REPLACE(target_to, '$', ''), ',', '') AS NUMERIC) AS target_to_num,
                CAST(REPLACE(REPLACE(target_from, '$', ''), ',', '') AS NUMERIC) AS target_from_num
            FROM item
        ) AS subquery
        WHERE target_to_num < target_from_num
        ORDER BY porcent DESC
        LIMIT 10;`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var item ChartItem
        // Escanear cada fila en la estructura `item`
        if err := rows.Scan(&item.Id, &item.Ticker, &item.Porcent); err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}

// func printBalances(conn *pgx.Conn) error {
//     rows, err := conn.Query(context.Background(), "SELECT id, balance FROM accounts")
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer rows.Close()
//     for rows.Next() {
//         var id uuid.UUID
//         var balance int
//         if err := rows.Scan(&id, &balance); err != nil {
//             log.Fatal(err)
//         }
//         log.Printf("%s: %d\n", id, balance)
//     }
//     return nil
// }

// func transferFunds(ctx context.Context, tx pgx.Tx, from uuid.UUID, to uuid.UUID, amount int) error {
//     // Read the balance.
//     var fromBalance int
//     if err := tx.QueryRow(ctx,
//         "SELECT balance FROM accounts WHERE id = $1", from).Scan(&fromBalance); err != nil {
//         return err
//     }

//     if fromBalance < amount {
//         log.Println("insufficient funds")
//     }

//     // Perform the transfer.
//     log.Printf("Transferring funds from account with ID %s to account with ID %s...", from, to)
//     if _, err := tx.Exec(ctx,
//         "UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, from); err != nil {
//         return err
//     }
//     if _, err := tx.Exec(ctx,
//         "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, to); err != nil {
//         return err
//     }
//     return nil
// }

// func deleteRows(ctx context.Context, tx pgx.Tx, one uuid.UUID, two uuid.UUID) error {
//     // Delete two rows into the "accounts" table.
//     log.Printf("Deleting rows with IDs %s and %s...", one, two)
//     if _, err := tx.Exec(ctx,
//         "DELETE FROM accounts WHERE id IN ($1, $2)", one, two); err != nil {
//         return err
//     }
//     return nil
// }

func Connet() {
    DB_URL := os.Getenv("DB_URL")
    // Read in connection string
    config, err := pgxpool.ParseConfig(DB_URL)
    if err != nil {
        log.Fatal(err)
    }

    config.MaxConns = 10;
    config.MinConns = 2;
    
    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        log.Fatal(err)
    }

    DB = pool

    // // Set up table
    // err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
    //     return initTable(context.Background(), tx)
    // })

    // // Insert initial rows
    // var accounts [4]uuid.UUID
    // for i := 0; i < len(accounts); i++ {
    //     accounts[i] = uuid.New()
    // }

    // err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
    //     return insertRows(context.Background(), tx, accounts)
    // })
    // if err == nil {
    //     log.Println("New rows created.")
    // } else {
    //     log.Fatal("error: ", err)
    // }

    // // Print out the balances
    // log.Println("Initial balances:")
    // printBalances(conn)

    // // Run a transfer
    // err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
    //     return transferFunds(context.Background(), tx, accounts[2], accounts[1], 100)
    // })
    // if err == nil {
    //     log.Println("Transfer successful.")
    // } else {
    //     log.Fatal("error: ", err)
    // }

    // // Print out the balances
    // log.Println("Balances after transfer:")
    // printBalances(conn)

    // // Delete rows
    // err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
    //     return deleteRows(context.Background(), tx, accounts[0], accounts[1])
    // })
    // if err == nil {
    //     log.Println("Rows deleted.")
    // } else {
    //     log.Fatal("error: ", err)
    // }

    // // Print out the balances
    // log.Println("Balances after deletion:")
    // printBalances(conn)
}

