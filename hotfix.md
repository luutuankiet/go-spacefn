 ## to get this working on chromebook del 3100 running lubuntu


## compile the bin
`go build`


## set your user in the input group

`sudo usermod -a -G input $USER`



## make sure /dev/uinput permission is fixed
/dev/uinput must have 666 with this 

- create this rule
```bash
sudo touch /etc/udev/rules.d/99-uinput.rules
```
- and content
```bash
KERNEL=="uinput", MODE="0666", GROUP="input"

``` 

- then run this

```bash
sudo udevadm control --reload-rules
sudo udevadm trigger

```



## to auto load

```bash
mkdir -p ~/.config/autostart
nano ~/.config/autostart/go-spacefn.desktop
```


`~/.congfig/autostart/spacefn.desktop`
```
[Desktop Entry]
Type=Application
Exec=/home/ken/spacefn.sh
Hidden=false
NoDisplay=false
X-GNOME-Autostart-enabled=true
Name=Go SpaceFN
Comment=Start Go SpaceFN on login
```


`~/.config/autostart/keyboard-settings.desktop`


```
[Desktop Entry]
Type=Application
Exec=xset r off
Hidden=false
NoDisplay=false
X-GNOME-Autostart-enabled=true
Name=Keyboard Settings
Comment=Set keyboard repeat rate

```


## to unload and config
`ps aux | grep spacefn`        
then pkill it