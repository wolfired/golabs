/**
 * tsc -m none -t ESNEXT gotv.ts && uglifyjs gotv.js -o gotv.js && move /Y gotv.js ..\testbed
 * cd ../gotv && tsc -m none -t ESNEXT gotv.ts && uglifyjs gunzip.min.js gotv.js -o gotv.js && mv gotv.js ../testbed && cd ../testbed
 */
document.body.style.margin = "0px";

const gotv = new class {
    private _ws: WebSocket;

    public constructor() {
    }

    public boot(): void {
        this._ws = new WebSocket("ws://{{.Host}}{{.Ws}}");
        this._ws.binaryType = "arraybuffer";
        this._ws.onerror = (ee: Event) => {
            console.log("error");
        };
        this._ws.onopen = (oe: Event) => {
            console.log("open");
            scrn.boot();
        };
        this._ws.onclose = (ce: CloseEvent) => {
            console.log("close");
        };
        this._ws.onmessage = (me: MessageEvent) => {
            console.log("message");
            let enable_zip = "{{.EnableZip}}";
            if ("true" === enable_zip) {
                var gunzip = new Zlib.Gunzip(new Uint8Array(me.data));
                var plain = gunzip.decompress();

                for (let i = 0; i < plain.length; ++i) {
                    scrn._cid.data[i] = plain[i];
                }
                scrn.draw();
            } else {
                let u8c = new Uint8ClampedArray(me.data);

                for (let i = 0; i < u8c.length; ++i) {
                    scrn._cid.data[i] = u8c[i];
                }
                scrn.draw();
            }



        };
    }

    public halt(): void {
        this._ws.close();
    }
}();

const gamepad = new class {
    private _gamepads: Gamepad[] = [];

    public constructor() {
    }

    public boot(): void {
        const hadGamepadEvent = "GamepadEvent" in window;
        const hadWebKitGamepadEvent = "WebKitGamepadEvent" in window;

        if (hadGamepadEvent) {
            window.addEventListener("gamepadconnected", (event: GamepadEvent): void => {
                this._gamepads[event.gamepad.index] = event.gamepad;
            });
            window.addEventListener("gamepaddisconnected", (event: GamepadEvent): void => {
                delete this._gamepads[event.gamepad.index];
            });
        } else if (hadWebKitGamepadEvent) {
            window.addEventListener("webkitgamepadconnected", (event: GamepadEvent): void => {
                this._gamepads[event.gamepad.index] = event.gamepad;
            });
            window.addEventListener("webkitgamepaddisconnected", (event: GamepadEvent): void => {
                delete this._gamepads[event.gamepad.index];
            });
        }

        const rAF = window.requestAnimationFrame || window["mozRequestAnimationFrame"] || window["webkitRequestAnimationFrame"];

        const fn = () => {
            this.scan();
            this.update();

            rAF(fn);
        };

        rAF(fn);
    }

    private scan(): void {
        const gamepads = void 0 !== navigator.getGamepads ? navigator.getGamepads() : (void 0 !== navigator["webkitGetGamepads"] ? navigator["webkitGetGamepads"]() : []);

        for (const gamepad of gamepads) {
            if (void 0 !== gamepad && null !== gamepad) {
                this._gamepads[gamepad.index] = gamepad;
            }
        }
    }

    private update(): void {
        let gp: number = -1.0;

        this._gamepads.forEach((gamepad: Gamepad) => {
            ++gp;

            this.debug("gamepad" + gp, "Gamepad: " + gamepad.id);

            let b: number = -1.0;
            gamepad.buttons.forEach((button: GamepadButton) => {
                ++b;

                this.debug("button" + b, "Button " + b + " pressed: " + button.pressed + ", value: " + button.value);
            });

            let a: number = -1.0;
            gamepad.axes.forEach((axis: number) => {
                ++a;

                this.debug("axis" + a, "Axis " + a + " value: " + axis.toFixed(4));
            });
        });
    }

    private debug(id: string, text: string): void {
        let div = document.getElementById(id) as HTMLDivElement;
        if (void 0 == div) {
            div = document.createElement("div");
            div.id = id;
            document.body.appendChild(div);
        }
        div.innerText = text;
    }
}();

const scrn = new class {
    private _can: HTMLCanvasElement;
    private _c2d: CanvasRenderingContext2D;
    public _cid: ImageData;

    public constructor() {
    }

    public boot(): void {
        this._can = document.createElement("canvas");

        document.body.appendChild(this._can);
        this._can.style.display = "block";

        this._c2d = this._can.getContext("2d");

        this.resize(parseInt("{{.Wid}}"), parseInt("{{.Hei}}"));
    }

    public resize(wid: number, hei: number): void {
        this._can.style.width = wid + "px";
        this._can.style.height = hei + "px";

        this._can.width = Math.round(this._can.clientWidth * window.devicePixelRatio);
        this._can.height = Math.round(this._can.clientHeight * window.devicePixelRatio);

        this._cid = this._c2d.getImageData(0, 0, wid, hei);
    }

    public draw(): void {
        this._c2d.putImageData(this._cid, 0, 0);
    }
}();
