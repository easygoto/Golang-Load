with open('./players.txt', 'w') as fp:
    fp.writelines(["login admin%d %d %d\n" % ((i + 1), i * 2, i * 3) for i in range(99999)])
