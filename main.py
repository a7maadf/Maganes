import re,sys, base64, json


print("Welcome to Maganes programming language, we fight feminists 24/7 here...\n")

try:
    x = sys.argv[1]
except Exception:
    print("No file specified")
    quit()



def lineArguments():
    try:
        f = open(sys.argv[1], "r")
    except Exception:
        print("Error, Check if file exists")
        quit()
        f = 2

    return f

def importingTheJson():
    f = open("cmd/cmds.json", "r")
    jsondic = json.load(f)
    return jsondic


def commandsCodes():
    f = lineArguments()
    jsondic = importingTheJson()
    TheWholeCode = ""  
    intolines = f.readlines()

    for i in intolines:
        line = list(i)
        fullline = ""
        for i in line:
            fullline += i
        for i in jsondic:
            if base64.b64decode(jsondic[i].encode("ascii")).decode("ascii") in fullline:
                print("Unknown", "'", base64.b64decode(jsondic[i].encode("ascii")).decode("ascii"),"'", "syntax found.\nPlease review your code...")
                quit()
            try:
                fullline = re.sub(i, base64.b64decode(jsondic[i].encode("ascii")).decode("ascii"), fullline)
            except Exception:
                pass
        TheWholeCode += fullline
    exec(TheWholeCode)





try:
    lineArguments()
    commandsCodes()
except Exception:
    print("Something doesn't seem right, please review your code and try again.")
