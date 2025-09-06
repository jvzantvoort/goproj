# structure

A templates are stored on ``~/.goproj/<name>`` with structure:

```
~/.goproj/<name>/
├── config.yml
├── tmux.tmpl
├── bash.tmpl
├── windows.tmpl
└── skel
    ├── linux
    │   └── bin
    │       └── command
    └── windows
        └── bin
            └── command.exe
```

# config file

Main config file (``config.yml``) contentse:

```yaml
---
template: <name>
pattern: ^.*$
version: 1
setup:
  clone:
    - name: projectname
      url: git@github.com:username/projectname.git
    - name: build
      url: git@github.com:username/buildsource.git
      destination: build
      branch: stable
  commands:
    - terraform do something cool
files:
  - name: tmux.tmpl
    destination: "{{.ConfigDir}}/{{.ProjectName}}.tmux"
    mode: 644
  - name: bash.tmpl
    destination: "{{.ConfigDir}}/{{.ProjectName}}.sh"
    mode: 644
  - name: windows.tmpl
    destination: "{{.ConfigDir}}/{{.ProjectName}}.bat"
```

## main

| parameter | description                            | default  |
|-----------|----------------------------------------|----------|
| template  | name of the template                   |          |
| pattern   | pattern for acceptance of new projects | ``^.*$`` |
| version   | version of the template                | 1        |
| setup     | actions to setup a new project         |          |
| files     | files to copy to                       |          |

## setup

| parameter | description |
|-----------|-------------|
| clone     | repositories to clone |
| commands  | commands to execute after cloning |

## files

| parameter | description | default |
|-----------|-------------|---------|
| clone | repositories to clone | |
| commands | commands to execute after cloning | |

# template files

TMUX template (``tmux.tmpl``) contense:
```
#  DESCRIPTION:  {{.ProjectDescription}}
#      WORKDIR:  {{.ProjectDir}}

set-option -g default-terminal "screen-256color"
set-option -g history-limit 10000
set-option -g monitor-activity on
set-option -g visual-activity off
set-option -sg escape-time 0
set-window-option -g xterm-keys on
set-window-option -g automatic-rename on

bind-key -r h select-pane -L  # move left
bind-key -r j select-pane -D  # move down
bind-key -r k select-pane -U  # move up
bind-key -r l select-pane -R  # move right

bind-key -r H resize-pane -L 2
bind-key -r J resize-pane -D 2
bind-key -r K resize-pane -U 2
bind-key -r L resize-pane -R 2

bind-key / command-prompt "split-window 'exec man %%'"

bind-key Tab last-window        # move to last active window

run-shell "bash {{.HomeDir}}/.bash/tmux.d/tmux_opt_source tmux-themepack/powerline/default/cyan.tmuxtheme"
```

```sh
#!/bin/bash

export PROJDIR="{{.ProjectDir}}"

if ! echo "$PATH" | grep -E -q "(^|:)${PROJDIR}/bin($|:)"
then
  export PATH="${PROJDIR}/bin:${PATH}"
fi

# Based on https://unix.stackexchange.com/questions/50208/how-to-change-the-working-directory-of-invoking-shell-using-a-script
function pcd()
{
  cd $PROJDIR/$1
}
```
