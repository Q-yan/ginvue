import sys

# 获取命令行参数列表
args = sys.argv

# 第一个参数是脚本名称，后续的参数是传递给脚本的参数
# 例如：python script.py argument1 argument2
# args[0] -> 'script.py'
# args[1] -> 'argument1'
# args[2] -> 'argument2'

# 处理命令行参数
if len(args) >= 3:
    arg1 = args[1]
    arg2 = args[2]
    print("Argument 1:", arg1)
    print("Argument 2:", arg2)
else:
    print("Please provide at least two arguments.")
