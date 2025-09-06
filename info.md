# Project Name: goproj

## Project Goals:

* Create a command-line tool in Go that serves as a backend for managing developer project environments.
* The tool should handle project locations, content, and version control sources (checkout, rebase, etc.).
* A key feature is the ability to generate a full project environment (e.g., a directory structure, cloned repositories) based on simple keywords (e.g., "Jira-12345").
* The tool should also provide a simple command to update all project sources.
* The ultimate goal is to pair with a frontend like tmux-project to create a seamless, automated development setup.

## Key Concepts and Functionality:

* Profiles: The tool should use configurable profiles (e.g., clientX, personal). These profiles will contain settings for source locations, project structures, and any specific commands.
* Keywords/Context: The user will provide a keyword (e.g., a ticket number, a client name) that the tool will use to select a profile and generate the environment.
* Source Management: The tool needs to be able to:
  * Check out source code from Git or other VCS repositories.
  * Automatically handle updates (e.g., pulling latest changes, rebasing branches).
* Configuration: The tool's behavior should be defined through a simple configuration file (e.g., YAML, TOML) that specifies profiles, project templates, and source locations.

## Development Assistance Focus:

When interacting with an AI assistant, I am looking for help with the following tasks:

1. Go Code Snippets: Provide idiomatic Go code for common tasks, such as:
   * Parsing command-line flags and arguments.
   * Reading and parsing a configuration file.
   * Executing shell commands (e.g., git clone, git rebase).
   * Managing file system operations (creating directories, copying files).
   * Handling concurrent operations for speed (e.g., updating multiple repositories in parallel).
2. Architectural Advice: Help with design decisions, such as:
   * Structuring the Go project (e.g., using packages for different functionalities like config, vcs, project).
   * Designing the configuration file schema.
   * Best practices for error handling and logging.
3. Specific Implementations: Provide guidance on implementing specific features, for example:
   * How to build a flexible system for defining and using project templates.
   * How to handle different types of sources (e.g., Git, Subversion) in a unified way.
   * How to create a robust and user-friendly CLI.
