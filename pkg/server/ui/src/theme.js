
export const theme = {
    "name": "d2",
    "rounding": 0,
    "spacing": 20,
    "defaultMode": "light",
    "global": {
      "fontFamily" : "Exocet Light",
      "colors": {
        "highlight": "#9d0208",
        "unique": "#C19965",
        "set": "#357E24",
        "rune": "#8C8C8D",
        "overall": "white",
        "brand": {
          "light": "#C19965",
          "dark": "#C19965"
        },
        "background": {
          "dark": "light-blue",
          "light": "#FFFFFF"
        },
        "background-back": {
          "dark": "dark-blue",
          "light": "#EEEEEE"
        },
        "background-front": {
          "dark": "lighter-blue",
          "light": "#FFFFFF"
        },
        "background-contrast": {
          "dark": "lighter-blue",
          "light": "#11111111"
        },
        "text": {
          "dark": "white",
          "light": "#333333"
        },
        "text-strong": {
          "dark": "#8C8C8D",
          "light": "#000000"
        },
        "text-weak": {
          "dark": "#8C8C8D",
          "light": "#444444"
        },
        "text-xweak": {
          "dark": "#616161",
          "light": "#666666"
        },
        "border": {
          "dark": "#8C8C8D",
          "light": "#8C8C8D"
        },
        "control": "brand",
        "active-background": "background-contrast",
        "active-text": "text-strong",
        "selected-background": "brand",
        "selected-text": "text-strong",
        "status-critical": "#FF4040",
        "status-warning": "#FFAA15",
        "status-ok": "#00C781",
        "status-unknown": "#CCCCCC",
        "status-disabled": "#CCCCCC",
        "graph-0": "brand",
        "graph-1": "status-warning",
        "dark-blue": {
          "dark": "#0D1318",
          "light": "#0D1318"
        },
        "dark-blue!": "",
        "blue": {
          "dark": "#0F1820",
          "light": "#0F1820"
        },
        "blue!": "",
        "lighter-blue": {
          "dark": "#141A20",
          "light": "#141A20"
        },
        "lighter-blue!": "",
        "light-blue": {
          "dark": "#111820",
          "light": "#111820"
        },
        "light-blue!": "",
        "focus": {
          "dark": "blue",
          "light": "blue"
        }
      },
      "hover" : {
        "shadow" : {
          "size" : "0px"
        }
      },
      "focus" : {
        "shadow" : {
          "size" : "0px"
        }
      },
      "font": {
        "family": "",
        "size": "15px",
        "height": "20px",
        "maxWidth": "300px"
      },
      "active": {
        "background": "active-background",
        "color": "active-text"
      },
      "hover": {
        "background": "active-background",
        "color": "active-text"
      },
      "selected": {
        "background": "selected-background",
        "color": "selected-text"
      },
      "control": {
        "border": {
          "radius": "0px"
        }
      },
      "drop": {
        "border": {
          "radius": "0px"
        }
      },
      "borderSize": {
        "xsmall": "1px",
        "small": "2px",
        "medium": "3px",
        "large": "10px",
        "xlarge": "20px"
      },
      "breakpoints": {
        "small": {
          "value": 720,
          "borderSize": {
            "xsmall": "1px",
            "small": "2px",
            "medium": "3px",
            "large": "5px",
            "xlarge": "10px"
          },
          "edgeSize": {
            "none": "0px",
            "hair": "1px",
            "xxsmall": "2px",
            "xsmall": "3px",
            "small": "5px",
            "medium": "10px",
            "large": "20px",
            "xlarge": "40px"
          },
          "size": {
            "xxsmall": "20px",
            "xsmall": "40px",
            "small": "80px",
            "medium": "160px",
            "large": "320px",
            "xlarge": "640px",
            "full": "100%"
          }
        },
        "medium": {
          "value": 1280
        },
        "large": {}
      },
      "edgeSize": {
        "none": "0px",
        "hair": "1px",
        "xxsmall": "3px",
        "xsmall": "5px",
        "small": "10px",
        "medium": "20px",
        "large": "40px",
        "xlarge": "80px",
        "responsiveBreakpoint": "small"
      },
      "input": {
        "padding": "10px",
        "weight": 600
      },
      "spacing": "20px",
      "size": {
        "xxsmall": "40px",
        "xsmall": "80px",
        "small": "160px",
        "medium": "320px",
        "large": "640px",
        "xlarge": "960px",
        "xxlarge": "1280px",
        "full": "100%"
      }
    },
    "tip": {
      "content": {
        "background": "background-contrast",
        "elevation": "0",
        "margin": "xsmall",
        "pad": {
          "vertical": "xsmall",
          "horizontal": "small"
        },
        "round": "0"
      }
    },
    "chart": {},
    "diagram": {
      "line": {}
    },
    "meter": {},
    "button": {
      "border": {
        "color" : "brand",
        "width": "1px",
        "radius": "0px"
      },
      "padding": {
        "vertical": "3px",
        "horizontal": "18px"
      },
      "hover" : {
        "shadow" : {
          "size" : "0px"
        }
      },
    },
    "checkBox": {
      "check": {
        "radius": "0px"
      },
      "toggle": {
        "radius": "20px",
        "size": "40px"
      },
      "size": "20px"
    },
    "radioButton": {
      "size": "20px"
    },
    "formField": {
      "border": {
        "color": "border",
        "error": {
          "color": {
            "dark": "white",
            "light": "status-critical"
          }
        },
        "position": "inner",
        "side": "bottom",
        "style": "solid"
      },
      "content": {
        "pad": "small"
      },
      "disabled": {
        "background": {
          "color": "status-disabled",
          "opacity": "medium"
        }
      },
      "error": {
        "color": "status-critical",
        "margin": {
          "vertical": "xsmall",
          "horizontal": "small"
        }
      },
      "help": {
        "color": "dark-3",
        "margin": {
          "start": "small"
        }
      },
      "info": {
        "color": "text-xweak",
        "margin": {
          "vertical": "xsmall",
          "horizontal": "small"
        }
      },
      "label": {
        "margin": {
          "vertical": "xsmall",
          "horizontal": "small"
        }
      },
      "margin": {
        "bottom": "small"
      },
      "round": "0px"
    },
    "calendar": {
      "small": {
        "fontSize": "11.666666666666666px",
        "lineHeight": 1.375,
        "daySize": "22.86px"
      },
      "medium": {
        "fontSize": "15px",
        "lineHeight": 1.45,
        "daySize": "45.71px"
      },
      "large": {
        "fontSize": "25px",
        "lineHeight": 1.11,
        "daySize": "91.43px"
      }
    },
    "clock": {
      "analog": {
        "hour": {
          "width": "7px",
          "size": "20px"
        },
        "minute": {
          "width": "3px",
          "size": "10px"
        },
        "second": {
          "width": "3px",
          "size": "8px"
        },
        "size": {
          "small": "60px",
          "medium": "80px",
          "large": "120px",
          "xlarge": "180px",
          "huge": "240px"
        }
      },
      "digital": {
        "text": {
          "xsmall": {
            "size": "8.333333333333332px",
            "height": 1.5
          },
          "small": {
            "size": "11.666666666666666px",
            "height": 1.43
          },
          "medium": {
            "size": "15px",
            "height": 1.375
          },
          "large": {
            "size": "18.333333333333332px",
            "height": 1.167
          },
          "xlarge": {
            "size": "21.666666666666668px",
            "height": 1.1875
          },
          "xxlarge": {
            "size": "28.333333333333336px",
            "height": 1.125
          }
        }
      }
    },
    "heading": {
      "level": {
        "1": {
          "small": {
            "size": "28px",
            "height": "33px",
            "maxWidth": "567px"
          },
          "medium": {
            "size": "42px",
            "height": "47px",
            "maxWidth": "833px"
          },
          "large": {
            "size": "68px",
            "height": "73px",
            "maxWidth": "1367px"
          },
          "xlarge": {
            "size": "95px",
            "height": "100px",
            "maxWidth": "1900px"
          }
        },
        "2": {
          "small": {
            "size": "25px",
            "height": "30px",
            "maxWidth": "500px"
          },
          "medium": {
            "size": "35px",
            "height": "40px",
            "maxWidth": "700px"
          },
          "large": {
            "size": "45px",
            "height": "50px",
            "maxWidth": "900px"
          },
          "xlarge": {
            "size": "55px",
            "height": "60px",
            "maxWidth": "1100px"
          }
        },
        "3": {
          "small": {
            "size": "22px",
            "height": "27px",
            "maxWidth": "433px"
          },
          "medium": {
            "size": "28px",
            "height": "33px",
            "maxWidth": "567px"
          },
          "large": {
            "size": "35px",
            "height": "40px",
            "maxWidth": "700px"
          },
          "xlarge": {
            "size": "42px",
            "height": "47px",
            "maxWidth": "833px"
          }
        },
        "4": {
          "small": {
            "size": "18px",
            "height": "23px",
            "maxWidth": "367px"
          },
          "medium": {
            "size": "22px",
            "height": "27px",
            "maxWidth": "433px"
          },
          "large": {
            "size": "25px",
            "height": "30px",
            "maxWidth": "500px"
          },
          "xlarge": {
            "size": "28px",
            "height": "33px",
            "maxWidth": "567px"
          }
        },
        "5": {
          "small": {
            "size": "13px",
            "height": "18px",
            "maxWidth": "267px"
          },
          "medium": {
            "size": "13px",
            "height": "18px",
            "maxWidth": "267px"
          },
          "large": {
            "size": "13px",
            "height": "18px",
            "maxWidth": "267px"
          },
          "xlarge": {
            "size": "13px",
            "height": "18px",
            "maxWidth": "267px"
          }
        },
        "6": {
          "small": {
            "size": "12px",
            "height": "17px",
            "maxWidth": "233px"
          },
          "medium": {
            "size": "12px",
            "height": "17px",
            "maxWidth": "233px"
          },
          "large": {
            "size": "12px",
            "height": "17px",
            "maxWidth": "233px"
          },
          "xlarge": {
            "size": "12px",
            "height": "17px",
            "maxWidth": "233px"
          }
        }
      }
    },
    "paragraph": {
      "small": {
        "size": "13px",
        "height": "18px",
        "maxWidth": "267px"
      },
      "medium": {
        "size": "15px",
        "height": "20px",
        "maxWidth": "300px"
      },
      "large": {
        "size": "18px",
        "height": "23px",
        "maxWidth": "367px"
      },
      "xlarge": {
        "size": "22px",
        "height": "27px",
        "maxWidth": "433px"
      },
      "xxlarge": {
        "size": "28px",
        "height": "33px",
        "maxWidth": "567px"
      }
    },
    "text": {
      "xsmall": {
        "size": "12px",
        "height": "17px",
        "maxWidth": "233px"
      },
      "small": {
        "size": "13px",
        "height": "18px",
        "maxWidth": "267px"
      },
      "medium": {
        "size": "15px",
        "height": "20px",
        "maxWidth": "300px"
      },
      "large": {
        "size": "18px",
        "height": "23px",
        "maxWidth": "367px"
      },
      "xlarge": {
        "size": "22px",
        "height": "27px",
        "maxWidth": "433px"
      },
      "xxlarge": {
        "size": "28px",
        "height": "33px",
        "maxWidth": "567px"
      }
    },
    "layer": {
      "background": {
        "dark": "#111111",
        "light": "#FFFFFF"
      }
    },
    "scale": 1
  }
  