import winreg
import os

LIB_DIR_Path = r"D:\codeSofts\ZeroMQ\CMake"

def createOpenVINODirPath():
    # 系统环境变量
    # parentKey = winreg.OpenKey(winreg.HKEY_LOCAL_MACHINE, r"\SYSTEM\CurrentControlSet\Control\Session Manager\Environment",
                        #  0, winreg.KEY_ALL_ACCESS | winreg.KEY_WOW64_64KEY)
    key=winreg.CreateKey(winreg.HKEY_LOCAL_MACHINE,r"SYSTEM\CurrentControlSet\Control\Session Manager\Environment")
    winreg.SetValueEx(key,"ZeroMQ_DIR",0,winreg.REG_SZ,LIB_DIR_Path)
    winreg.CloseKey(key)
    print("create dir done")
    return

if __name__=="__main__":
    if(not os.path.exists(LIB_DIR_Path)):
        # 如果不存在
        print(LIB_DIR_Path," not exists")
    else:
        createOpenVINODirPath()
