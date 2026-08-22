# xpense tracker 

`budgeter` is a Go CLI for loading statements, loading transactions for a statement, and searching statements with a Bubble Tea terminal view.

## Build

```sh
go build ./cmd/budgeter
```

## Commands

```sh
budgeter [--db path] statement load
budgeter [--db path] statement list
budgeter [--db path] statement delete
budgeter [--db path] statement tx load
budgeter [--db path] statement search
budgeter [--db path] statement show
budgeter [--db path] statement tx
```

The default database path is `~/.budgeter/budgeter.db`.

## Examples

```sh
budgeter statement load --date 2026-08-20 --account cheque --balance 12000.00 --income 5000.00 --costs 2000.00 --earnings 3000.00 --savings 1000.00
budgeter statement tx load --statement-id 1 --date 2026-08-20 --bucket cost --category Groceries --description "Market" --amount -42.10
budgeter statement list --limit 25
budgeter statement delete --id 1 --yes
budgeter statement search --account cheque --from 2026-08-01 --top-10
budgeter statement tx --statement-id 1 --bucket cost --category Groceries --all
```

Statement loads use the latest saved statement as defaults. If you load a new statement and omit `account`, `income`, `costs`, `earnings`, or `savings`, those values are copied from the latest statement in the database. A load is rejected when a statement already exists for that statement's month.

Statement deletes require `--yes` and remove the statement plus its transactions.

## CSV

Statement CSV headers:

```csv
date,account,balance,income,costs,earnings,savings,notes
```

Transaction CSV headers:

```csv
date,bucket,category,desc,amount
```

Sample transaction CSV:

```csv
date,bucket,category,desc,amount
2026-08-01,income,Miscellaneous,Monthly salary,5000.00
2026-08-02,cost,Groceries,Market,-42.10
2026-08-03,cost,Transport,Fuel,-65.50
2026-08-04,savings,Investments,Savings transfer,-500.00
```

Load it with:

```sh
budgeter statement tx load --statement-id 1 --file txs.csv
```

Transaction buckets are `income`, `cost`, and `savings`. Bucketed transactions are used to calculate statement `income`, `costs`, `earnings`, `savings`, and `balance` when a statement is shown or listed. `earnings` is `income - costs`; `balance` is `income - costs - savings`. Existing unbucketed transactions are still displayed, but they are not included in calculated statement rows.

Transaction categories are `Bills`, `Utilities`, `Groceries`, `Eating Out`, `Subscriptions`, `Transport`, `Shopping`, `Entertainment`, `Debt Repayments`, `Investments`, `Insurance`, `Healthcare`, `Personal Care`, and `Miscellaneous`.

Transaction lists can be filtered by bucket and category:

```sh
budgeter statement tx --statement-id 1 --bucket cost --category Groceries --all
```

Transaction read flags are mutually exclusive:

```sh
--all
--top-10
--top-100
--current-month
```
