
# Allyas Manager

Alias Manager is a simple CLI tool written in Go that helps you manage your shell aliases. It allows you to store aliases in an `ally_aliases` file and source them into your default shell configuration file. The tool automatically detects the shell you're using by checking the `$SHELL` environment variable.

## Installation

### Requirements

- [Go](https://golang.org/) (Go 1.19 or higher)
- [Make](https://www.gnu.org/software/make/)

### Steps to Install

1. Clone the repository:

   ```bash
   git clone https://github.com/prettyletto/allyas.git
   cd alias-manager
   ```

2. Build the tool using `make`:

   ```bash
   make
   ```

3. (Optional) To install the tool globally, move the compiled binary to a directory in your `PATH`:

   ```bash
   sudo mv allyas /usr/local/bin/
   ```

Now you can use `allyas` from any directory.

## Usage

The tool provides several commands. Here's how to use them:

### `create`

The `create` command allows you to add a new alias.

**Usage:**

```bash
allyas create [alias_name] "[command]" "#[description]"
```

- **`alias_name`**: The name of the alias you want to create.
- **`command`**: The command the alias will execute.
- **`description`** (optional): A comment describing the alias. If no description is provided, the tool will automatically generate one based on the alias name.

**Example:**

```bash
allyas create your_alias "your_command" "#description"
```

This will create the following entry in the `ally_aliases` file:

```bash
#description
alias your_alias="your_command"
```

**Behavior without a description:**
If you don't provide a description, the tool will generate a default description based on the alias name. For example:

```bash
allyas create my_alias "ls -la"
```

This will produce:

```bash
#my_alias
alias my_alias="ls -la"
```

### `list`

The `list` command shows all the aliases in the `ally_aliases` file.

**Usage:**

```bash
allyas list
```

This will display all the aliases currently stored in the `ally_aliases` file, along with their descriptions.

**Example Output:**

```bash
#My_alias 
my_alias  ->  my_command  #my_description

#another_alias
another_alias -> another_command #another_description
```

Here's an improved and clearer version of your `edit` command section:

---

### `edit`

The `edit` command allows you to modify an existing alias.

**Usage:**

```bash
allyas edit [old_name] [option] [new_value]
```

- **`old_name`**: The current name of the alias you want to modify.
- **`option`**: The option you want to modify:
  - **`-a`**: Edit the alias name.
  - **`-c`**: Edit the command associated with the alias.
  - **`-d`**: Edit the description for the alias.
- **`new_value`**: The new value to replace the current one. This could be a new alias name, command, or description.

**Examples:**

1. **Edit the alias name:**

```bash
allyas edit my_alias -a new_alias_name
```

This will change the alias from `my_alias` to `new_alias_name` in the `ally_aliases` file.

2. **Edit the command:**

```bash
allyas edit my_alias -c "ls -alh"
```

This will update the command for `my_alias` to `ls -alh`.

3. **Edit the description:**

```bash
allyas edit my_alias -d "List files in human-readable format"
```

This will change the description of `my_alias` to `"List files in human-readable format"`, resulting in the following entry in the `ally_aliases` file:

```bash
# List files in human-readable format
alias my_alias="ls -alh"
```

---

### `remove`

The `remove` command allows you to delete an alias.

**Usage:**

```bash
allyas remove [alias_name]
```

- **`alias_name`**: The name of the alias you want to remove.

**Example:**

```bash
allyas remove my_alias
```

This will delete the `my_alias` entry from the `ally_aliases` file.

---






Now based on these examples and commands i need to build a -h for my application 
