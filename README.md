# Do you wanna list watchers of all your projects?

I don't know why github do not send notifications about new watchers, but I
think, repository watchers it's more importantly than stargazers, and I really
want see, who follow my projects.

## Installation

Project written in go, so for install you should just write:

```
go get github.com/kovetskiy/github-watchers
```

and **github-watchers** will be installed.

## Usage

Synopsis:
```
github-watchers <username> [<repository>...]
```

How you can see, you should specify your username, and if you have too much
repositories and do not want see all watchers to all repositories, you can
specify needed repositories, if not, than **gihub-watchers** will get a list of
all public repositories of specified user.

### Output

```
https://github.com/kovetskiy/dotfiles:
kovetskiy   https://github.com/kovetskiy
seletskiy   https://github.com/seletskiy
brnv        https://github.com/brnv
ARoiD       https://github.com/ARoiD

https://github.com/kovetskiy/zsh-fastcd:
kovetskiy   https://github.com/kovetskiy
seletskiy   https://github.com/seletskiy
brnv        https://github.com/brnv

https://github.com/kovetskiy/zsh-add-params:
kovetskiy   https://github.com/kovetskiy
seletskiy   https://github.com/seletskiy

https://github.com/kovetskiy/mcabber-focus:
kovetskiy   https://github.com/kovetskiy
seletskiy   https://github.com/seletskiy

https://github.com/kovetskiy/scut:
kovetskiy     https://github.com/kovetskiy
seletskiy     https://github.com/seletskiy
romanyellow   https://github.com/romanyellow
anoland       https://github.com/anoland
grimnight     https://github.com/grimnight
```
