# PNG - Project Number Generator

PNG is a CLI program created to generate unique IDs for Harmony's project report sheets. It allows you to manage teams, work types, and projects by adding and listing them in a SQLite database.

## Installation

### Windows

1. Build from the source or download the `png.exe` file from the [releases page](~https://github.com/torives/png/releases~).
2. Create a `png` directory in `C:\Program Files` (or any other directory) and place `png.exe` there.

## Usage

The first step is to open a terminal window in the directory where `png.exe` is located. Then you can use PNG via an **interactive prompt** or invoking **commands**. Both options are functionally the same, aside from the fact you can't [specify a database](#database) when using the interactive prompt.

To start the interactive prompt, run PNG without parameters:

```bash
./png
```

Use the arrow keys to navigate the options and press enter to select one.

Alternatively, execute a command directly. For example:

```bash
./png team list
```

The [Features](#features) section will delve deeper into the options available in both cases.

## Features

### Add a project

You can add a new project by providing a team and work type, and PNG will generate and return its unique ID. All IDs are formatted as: `AAA-BB-#`, where:

- `AAA` is a three-uppercase letter code identifying the team responsible for the project
- `BB` is a two-uppercase letter code identifying the type of work
- `#` is a unique number

#### Examples:

**Interactive prompt:**

1. Select `add a new project`.
2. Type the team's name and press enter.
3. Type the work type and press enter.
4. The project is created and its ID is returned.

> ⚠️ Notice that the CLI won't accept inputs with the wrong format (team names with four letters), but it will accept teams and work types that are not in the database (it will return a "not found" error).

**CLI:**

```bash
./png project add -t ABC -w XY
```

### List all projects

Use this command to list all projects in no particular order.

#### Examples:

**Interactive prompt:**

1. Select `list all projects`.
2. PNG returns a list of all project IDs.

**CLI:**

```bash
./png project list
```

### Add a new team

Use this command to add a new team. Make sure its name contains only three uppercase letters.

#### Examples:

**Interactive prompt:**

1. Select `add a new team`.
2. Type the team's name and press enter.
3. PNG returns a success message.

**CLI:**

```bash
./png team add ABC
```

### List all teams

Use this command to list all available teams in no particular order.

#### Examples:

**Interactive prompt:**

1. Select `list all teams`.
2. PNG returns a list of all teams.

**CLI:**

```bash
./png team list
```

### Add a new work type

Use this command to add a new work type. Make sure its name contains only two uppercase letters.

#### Examples:

**Interactive prompt:**

1. Select `add a new work type`.
2. Type the work type’s name and press enter.
3. PNG returns a success message.

**CLI:**

```bash
./png worktype add XY
```

### List all work types

Use this command to list all available work types in no particular order.

**Interactive prompt:**

1. Select `list all work types`.
2. PNG returns a list of all work types.

**CLI:**

```bash
./png worktype list
```

## Database

By default, the PNG CLI uses a SQLite database file named `png.sqlite`. This file will be automatically created in the same directory the executable is located and should be backed up periodically to avoid data loss.

For convenience, the database comes preloaded with the following teams and work types:

- Teams
  - FOR (Formulation) – Everything related to the assembly of the final product.
  - ANA (Analytical) – Characterization of ingredients, proteins, and evaluation of processing eﬀects on ingredients.
  - MIC (Microbiology) – Everything related to microbiota growth response.
  - PRO (Protein) – From strain development to production, not including characterization.
- Types of work
  - MA (Mechanism of Action).
  - ES (Emulsion Studies).
  - IC (Ingredient Characterization).
  - PT (Prototypes).
  - PP (Production Process).

Optionally, it is possible to execute commands in a different database file using the `--db` flag.

Example:

```bash
./png project list --db /path/to/another/database.sqlite
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.