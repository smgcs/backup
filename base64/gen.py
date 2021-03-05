#!/usr/bin/python
import base64

def main():
    inp = raw_input("input src:")
    out = base64.b64encode(inp)
    print out
    print out[::5]
if __name__ == "__main__":
    main()
