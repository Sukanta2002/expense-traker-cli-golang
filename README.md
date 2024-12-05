# Expense Tracker

A simple command line tool to track your expenses.

## Features

- Add expenses
- View expenses
- Delete expenses
- View summary of expenses

## TODO

- Add expense categories and allow users to filter expenses by category.
- Allow users to set a budget for each month and show a warning when the user exceeds the budget.
- Allow users to export expenses to a CSV file.

## Usage

### Add an expense

```bash
expense-traker-cli add -d "dth" -a 100 # or expense-traker-cli add --description "dth" --amount 100
```

### View expenses

```bash
expense-traker-cli view
```

### Delete an expense

```bash
expense-traker-cli delete -i 1 # or expense-traker-cli delete --id 1
```

### View summary of expenses

```bash
expense-traker-cli summary
```
