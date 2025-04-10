package main

import evdev "github.com/egorse/golang-evdev"

var keyMap1 = map[int]int{
	// arrows
	evdev.KEY_K: evdev.KEY_UP,
	evdev.KEY_H: evdev.KEY_LEFT,
	evdev.KEY_J: evdev.KEY_DOWN,
	evdev.KEY_L: evdev.KEY_RIGHT,

	// pgup/pgdown
	// evdev.KEY_R: evdev.KEY_PAGEUP,
	// evdev.KEY_F: evdev.KEY_PAGEDOWN,

	// home/end
	evdev.KEY_Y: evdev.KEY_HOME,
	evdev.KEY_O: evdev.KEY_END,

	// Fx keys
	// evdev.KEY_BACK: evdev.KEY_F1,
	// evdev.KEY_FORWARD: evdev.KEY_F2,
	// evdev.KEY_REFRESH: evdev.KEY_F3,
	// // evdev.KEY_4: evdev.KEY_F4,
	// evdev.KEY_SCALE: evdev.KEY_F5,
	// evdev.KEY_BRIGHTNESSDOWN: evdev.KEY_F6,
	// evdev.KEY_BRIGHTNESSUP: evdev.KEY_F7,
	// evdev.KEY_MUTE: evdev.KEY_F8,
	// evdev.KEY_VOLUMEDOWN: evdev.KEY_F9,
	// evdev.KEY_VOLUMEUP: evdev.KEY_F10,

	// Esc and `
	evdev.KEY_N:   evdev.KEY_GRAVE,
	evdev.KEY_M:   evdev.KEY_SLASH,
	evdev.KEY_GRAVE: evdev.KEY_ESC,
	
	// other
	evdev.KEY_D: evdev.KEY_BACKSPACE,
	evdev.KEY_B: evdev.KEY_ENTER,
	// evdev.KEY_LEFTCTRL: evdev.KEY_LEFTALT,
	// evdev.KEY_LEFTALT: evdev.KEY_LEFTCTRL,
}
