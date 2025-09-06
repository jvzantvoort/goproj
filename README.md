[![forthebadge](https://forthebadge.com/images/badges/made-with-crayons.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/designed-in-etch-a-sketch.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/you-didnt-ask-for-this.svg)](https://forthebadge.com)

# goproj

**goproj** creates, maintains, archives and removes projects.

`goproj` is a command-line tool designed to streamline and automate
the setup and management of developer environments. By providing a
flexible and configurable backend, it aims to simplify a developer's
workflow, making it easy to jump into a new project or update an
existing one with minimal effort.

This tool is intended to serve as the backend for `tmux-project`,
providing the core logic for managing project locations, content,
and source repositories.

## Intentions

The primary goal of `goproj` is to reduce the friction associated
with setting up a new development environment. Often, developers
need to handle multiple clients, projects, and repositories, each
with its own specific setup requirements. `goproj` addresses this
by:

* **Centralizing Project Management:** It manages project locations
  and their associated content, providing a single source of truth
  for all your projects.
* **Automating Source Management:** It handles the entire lifecycle
  of source code, from initial checkout to routine updates (e.g.,
  `rebase`). This ensures that your working directories are always
  up-to-date with a simple command.
* **Enabling Context-Based Generation:** A core feature is the
  ability to generate a complete project environment based on a
  simple keyword, such as a Jira ticket number or an email subject.
  This allows for the creation of a tailored workspace based on the
  immediate task at hand.
* **Facilitating Consistency:** By using configurable profiles,
  `goproj` helps maintain consistent development environments across
  different projects and team members.

## Design Notions

The design of `goproj` is guided by the following principles:

* **Modularity:** The tool will be built with a modular
  architecture, allowing for easy extension and integration with new
  sources, VCS systems, or project types.
* **Configuration over Convention:** While it will provide sensible
  defaults, the tool's behavior will be highly configurable. Users
  should be able to define their own profiles, source locations, and
  command hooks.
* **Idempotency:** Operations like "generate" or "update" should be
  repeatable without causing unintended side effects. Running a
  command a second time should produce the same result as the first,
  if the state hasn't changed.
* **User-Centric CLI:** The command-line interface will be intuitive
  and easy to use, providing clear feedback and sensible defaults
  while still allowing for detailed customization.
* **API-Driven Backend:** The core functionality will be exposed
  through a clear API, allowing `tmux-project` and other potential
  frontends to interact with it seamlessly. This separation of
  concerns ensures that the backend remains robust and independent
  of any specific frontend.

This design ensures that `goproj` is not just a tool for personal
use but a solid foundation that can be used by other tools and in
team environments.
