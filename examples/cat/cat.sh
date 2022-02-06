#!/bin/sh
go run .
montage -geometry 256x256+8+8 Cat-main-main--000.tmp.png Cat-main-main--000.debug.png Cat-main-main--000.debug.png
convert Cat-main-main--000.debug.png -bordercolor white -border 8 Cat-main-main--000.debug.png
montage -geometry 256x256+8+8 Cat-*-main--000.tmp.png montage_000.png
montage -geometry 256x256+8+8 Cat-main-main--001.tmp.png Cat-main-main--001.debug.png Cat-main-main--001.debug.png
convert Cat-main-main--001.debug.png -bordercolor white -border 8 Cat-main-main--001.debug.png
montage -geometry 256x256+8+8 Cat-*-main--001.tmp.png montage_001.png
montage -geometry 256x256+8+8 Cat-main-main--002.tmp.png Cat-main-main--002.debug.png Cat-main-main--002.debug.png
convert Cat-main-main--002.debug.png -bordercolor white -border 8 Cat-main-main--002.debug.png
montage -geometry 256x256+8+8 Cat-*-main--002.tmp.png montage_002.png
montage -geometry 256x256+8+8 Cat-main-main--003.tmp.png Cat-main-main--003.debug.png Cat-main-main--003.debug.png
convert Cat-main-main--003.debug.png -bordercolor white -border 8 Cat-main-main--003.debug.png
montage -geometry 256x256+8+8 Cat-*-main--003.tmp.png montage_003.png
montage -geometry 256x256+8+8 Cat-main-main--004.tmp.png Cat-main-main--004.debug.png Cat-main-main--004.debug.png
convert Cat-main-main--004.debug.png -bordercolor white -border 8 Cat-main-main--004.debug.png
montage -geometry 256x256+8+8 Cat-*-main--004.tmp.png montage_004.png
montage -geometry 256x256+8+8 Cat-main-main--005.tmp.png Cat-main-main--005.debug.png Cat-main-main--005.debug.png
convert Cat-main-main--005.debug.png -bordercolor white -border 8 Cat-main-main--005.debug.png
montage -geometry 256x256+8+8 Cat-*-main--005.tmp.png montage_005.png
montage -geometry 256x256+8+8 Cat-main-main--006.tmp.png Cat-main-main--006.debug.png Cat-main-main--006.debug.png
convert Cat-main-main--006.debug.png -bordercolor white -border 8 Cat-main-main--006.debug.png
montage -geometry 256x256+8+8 Cat-*-main--006.tmp.png montage_006.png
montage -geometry 256x256+8+8 Cat-main-main--007.tmp.png Cat-main-main--007.debug.png Cat-main-main--007.debug.png
convert Cat-main-main--007.debug.png -bordercolor white -border 8 Cat-main-main--007.debug.png
montage -geometry 256x256+8+8 Cat-*-main--007.tmp.png montage_007.png
montage -geometry 256x256+8+8 Cat-main-main--008.tmp.png Cat-main-main--008.debug.png Cat-main-main--008.debug.png
convert Cat-main-main--008.debug.png -bordercolor white -border 8 Cat-main-main--008.debug.png
montage -geometry 256x256+8+8 Cat-*-main--008.tmp.png montage_008.png
montage -geometry 256x256+8+8 Cat-main-main--009.tmp.png Cat-main-main--009.debug.png Cat-main-main--009.debug.png
convert Cat-main-main--009.debug.png -bordercolor white -border 8 Cat-main-main--009.debug.png
montage -geometry 256x256+8+8 Cat-*-main--009.tmp.png montage_009.png
montage -geometry 256x256+8+8 Cat-main-main--010.tmp.png Cat-main-main--010.debug.png Cat-main-main--010.debug.png
convert Cat-main-main--010.debug.png -bordercolor white -border 8 Cat-main-main--010.debug.png
montage -geometry 256x256+8+8 Cat-*-main--010.tmp.png montage_010.png
montage -geometry 256x256+8+8 Cat-main-main--011.tmp.png Cat-main-main--011.debug.png Cat-main-main--011.debug.png
convert Cat-main-main--011.debug.png -bordercolor white -border 8 Cat-main-main--011.debug.png
montage -geometry 256x256+8+8 Cat-*-main--011.tmp.png montage_011.png
montage -geometry 256x256+8+8 Cat-main-main--012.tmp.png Cat-main-main--012.debug.png Cat-main-main--012.debug.png
convert Cat-main-main--012.debug.png -bordercolor white -border 8 Cat-main-main--012.debug.png
montage -geometry 256x256+8+8 Cat-*-main--012.tmp.png montage_012.png
montage -geometry 256x256+8+8 Cat-main-main--013.tmp.png Cat-main-main--013.debug.png Cat-main-main--013.debug.png
convert Cat-main-main--013.debug.png -bordercolor white -border 8 Cat-main-main--013.debug.png
montage -geometry 256x256+8+8 Cat-*-main--013.tmp.png montage_013.png
montage -geometry 256x256+8+8 Cat-main-main--014.tmp.png Cat-main-main--014.debug.png Cat-main-main--014.debug.png
convert Cat-main-main--014.debug.png -bordercolor white -border 8 Cat-main-main--014.debug.png
montage -geometry 256x256+8+8 Cat-*-main--014.tmp.png montage_014.png
montage -geometry 256x256+8+8 Cat-main-main--015.tmp.png Cat-main-main--015.debug.png Cat-main-main--015.debug.png
convert Cat-main-main--015.debug.png -bordercolor white -border 8 Cat-main-main--015.debug.png
montage -geometry 256x256+8+8 Cat-*-main--015.tmp.png montage_015.png
convert Cat-main-main--*.tmp.png Cat_main_main.gif
convert Cat-main-main--*.debug.png Cat_main_main_debug.gif
montage -geometry 256x256+8+8 Cat-han-main--000.tmp.png Cat-han-main--000.debug.png Cat-han-main--000.debug.png
convert Cat-han-main--000.debug.png -bordercolor white -border 8 Cat-han-main--000.debug.png
montage -geometry 256x256+8+8 Cat-*-main--000.tmp.png montage_000.png
montage -geometry 256x256+8+8 Cat-han-main--001.tmp.png Cat-han-main--001.debug.png Cat-han-main--001.debug.png
convert Cat-han-main--001.debug.png -bordercolor white -border 8 Cat-han-main--001.debug.png
montage -geometry 256x256+8+8 Cat-*-main--001.tmp.png montage_001.png
montage -geometry 256x256+8+8 Cat-han-main--002.tmp.png Cat-han-main--002.debug.png Cat-han-main--002.debug.png
convert Cat-han-main--002.debug.png -bordercolor white -border 8 Cat-han-main--002.debug.png
montage -geometry 256x256+8+8 Cat-*-main--002.tmp.png montage_002.png
montage -geometry 256x256+8+8 Cat-han-main--003.tmp.png Cat-han-main--003.debug.png Cat-han-main--003.debug.png
convert Cat-han-main--003.debug.png -bordercolor white -border 8 Cat-han-main--003.debug.png
montage -geometry 256x256+8+8 Cat-*-main--003.tmp.png montage_003.png
montage -geometry 256x256+8+8 Cat-han-main--004.tmp.png Cat-han-main--004.debug.png Cat-han-main--004.debug.png
convert Cat-han-main--004.debug.png -bordercolor white -border 8 Cat-han-main--004.debug.png
montage -geometry 256x256+8+8 Cat-*-main--004.tmp.png montage_004.png
montage -geometry 256x256+8+8 Cat-han-main--005.tmp.png Cat-han-main--005.debug.png Cat-han-main--005.debug.png
convert Cat-han-main--005.debug.png -bordercolor white -border 8 Cat-han-main--005.debug.png
montage -geometry 256x256+8+8 Cat-*-main--005.tmp.png montage_005.png
montage -geometry 256x256+8+8 Cat-han-main--006.tmp.png Cat-han-main--006.debug.png Cat-han-main--006.debug.png
convert Cat-han-main--006.debug.png -bordercolor white -border 8 Cat-han-main--006.debug.png
montage -geometry 256x256+8+8 Cat-*-main--006.tmp.png montage_006.png
montage -geometry 256x256+8+8 Cat-han-main--007.tmp.png Cat-han-main--007.debug.png Cat-han-main--007.debug.png
convert Cat-han-main--007.debug.png -bordercolor white -border 8 Cat-han-main--007.debug.png
montage -geometry 256x256+8+8 Cat-*-main--007.tmp.png montage_007.png
montage -geometry 256x256+8+8 Cat-han-main--008.tmp.png Cat-han-main--008.debug.png Cat-han-main--008.debug.png
convert Cat-han-main--008.debug.png -bordercolor white -border 8 Cat-han-main--008.debug.png
montage -geometry 256x256+8+8 Cat-*-main--008.tmp.png montage_008.png
montage -geometry 256x256+8+8 Cat-han-main--009.tmp.png Cat-han-main--009.debug.png Cat-han-main--009.debug.png
convert Cat-han-main--009.debug.png -bordercolor white -border 8 Cat-han-main--009.debug.png
montage -geometry 256x256+8+8 Cat-*-main--009.tmp.png montage_009.png
montage -geometry 256x256+8+8 Cat-han-main--010.tmp.png Cat-han-main--010.debug.png Cat-han-main--010.debug.png
convert Cat-han-main--010.debug.png -bordercolor white -border 8 Cat-han-main--010.debug.png
montage -geometry 256x256+8+8 Cat-*-main--010.tmp.png montage_010.png
montage -geometry 256x256+8+8 Cat-han-main--011.tmp.png Cat-han-main--011.debug.png Cat-han-main--011.debug.png
convert Cat-han-main--011.debug.png -bordercolor white -border 8 Cat-han-main--011.debug.png
montage -geometry 256x256+8+8 Cat-*-main--011.tmp.png montage_011.png
montage -geometry 256x256+8+8 Cat-han-main--012.tmp.png Cat-han-main--012.debug.png Cat-han-main--012.debug.png
convert Cat-han-main--012.debug.png -bordercolor white -border 8 Cat-han-main--012.debug.png
montage -geometry 256x256+8+8 Cat-*-main--012.tmp.png montage_012.png
montage -geometry 256x256+8+8 Cat-han-main--013.tmp.png Cat-han-main--013.debug.png Cat-han-main--013.debug.png
convert Cat-han-main--013.debug.png -bordercolor white -border 8 Cat-han-main--013.debug.png
montage -geometry 256x256+8+8 Cat-*-main--013.tmp.png montage_013.png
montage -geometry 256x256+8+8 Cat-han-main--014.tmp.png Cat-han-main--014.debug.png Cat-han-main--014.debug.png
convert Cat-han-main--014.debug.png -bordercolor white -border 8 Cat-han-main--014.debug.png
montage -geometry 256x256+8+8 Cat-*-main--014.tmp.png montage_014.png
montage -geometry 256x256+8+8 Cat-han-main--015.tmp.png Cat-han-main--015.debug.png Cat-han-main--015.debug.png
convert Cat-han-main--015.debug.png -bordercolor white -border 8 Cat-han-main--015.debug.png
montage -geometry 256x256+8+8 Cat-*-main--015.tmp.png montage_015.png
convert Cat-han-main--*.tmp.png Cat_han_main.gif
convert Cat-han-main--*.debug.png Cat_han_main_debug.gif
montage -geometry 256x256+8+8 Cat-cat-cat-main--000.tmp.png Cat-cat-cat-main--000.debug.png Cat-cat-cat-main--000.debug.png
convert Cat-cat-cat-main--000.debug.png -bordercolor white -border 8 Cat-cat-cat-main--000.debug.png
montage -geometry 256x256+8+8 Cat-*-main--000.tmp.png montage_000.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--001.tmp.png Cat-cat-cat-main--001.debug.png Cat-cat-cat-main--001.debug.png
convert Cat-cat-cat-main--001.debug.png -bordercolor white -border 8 Cat-cat-cat-main--001.debug.png
montage -geometry 256x256+8+8 Cat-*-main--001.tmp.png montage_001.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--002.tmp.png Cat-cat-cat-main--002.debug.png Cat-cat-cat-main--002.debug.png
convert Cat-cat-cat-main--002.debug.png -bordercolor white -border 8 Cat-cat-cat-main--002.debug.png
montage -geometry 256x256+8+8 Cat-*-main--002.tmp.png montage_002.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--003.tmp.png Cat-cat-cat-main--003.debug.png Cat-cat-cat-main--003.debug.png
convert Cat-cat-cat-main--003.debug.png -bordercolor white -border 8 Cat-cat-cat-main--003.debug.png
montage -geometry 256x256+8+8 Cat-*-main--003.tmp.png montage_003.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--004.tmp.png Cat-cat-cat-main--004.debug.png Cat-cat-cat-main--004.debug.png
convert Cat-cat-cat-main--004.debug.png -bordercolor white -border 8 Cat-cat-cat-main--004.debug.png
montage -geometry 256x256+8+8 Cat-*-main--004.tmp.png montage_004.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--005.tmp.png Cat-cat-cat-main--005.debug.png Cat-cat-cat-main--005.debug.png
convert Cat-cat-cat-main--005.debug.png -bordercolor white -border 8 Cat-cat-cat-main--005.debug.png
montage -geometry 256x256+8+8 Cat-*-main--005.tmp.png montage_005.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--006.tmp.png Cat-cat-cat-main--006.debug.png Cat-cat-cat-main--006.debug.png
convert Cat-cat-cat-main--006.debug.png -bordercolor white -border 8 Cat-cat-cat-main--006.debug.png
montage -geometry 256x256+8+8 Cat-*-main--006.tmp.png montage_006.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--007.tmp.png Cat-cat-cat-main--007.debug.png Cat-cat-cat-main--007.debug.png
convert Cat-cat-cat-main--007.debug.png -bordercolor white -border 8 Cat-cat-cat-main--007.debug.png
montage -geometry 256x256+8+8 Cat-*-main--007.tmp.png montage_007.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--008.tmp.png Cat-cat-cat-main--008.debug.png Cat-cat-cat-main--008.debug.png
convert Cat-cat-cat-main--008.debug.png -bordercolor white -border 8 Cat-cat-cat-main--008.debug.png
montage -geometry 256x256+8+8 Cat-*-main--008.tmp.png montage_008.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--009.tmp.png Cat-cat-cat-main--009.debug.png Cat-cat-cat-main--009.debug.png
convert Cat-cat-cat-main--009.debug.png -bordercolor white -border 8 Cat-cat-cat-main--009.debug.png
montage -geometry 256x256+8+8 Cat-*-main--009.tmp.png montage_009.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--010.tmp.png Cat-cat-cat-main--010.debug.png Cat-cat-cat-main--010.debug.png
convert Cat-cat-cat-main--010.debug.png -bordercolor white -border 8 Cat-cat-cat-main--010.debug.png
montage -geometry 256x256+8+8 Cat-*-main--010.tmp.png montage_010.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--011.tmp.png Cat-cat-cat-main--011.debug.png Cat-cat-cat-main--011.debug.png
convert Cat-cat-cat-main--011.debug.png -bordercolor white -border 8 Cat-cat-cat-main--011.debug.png
montage -geometry 256x256+8+8 Cat-*-main--011.tmp.png montage_011.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--012.tmp.png Cat-cat-cat-main--012.debug.png Cat-cat-cat-main--012.debug.png
convert Cat-cat-cat-main--012.debug.png -bordercolor white -border 8 Cat-cat-cat-main--012.debug.png
montage -geometry 256x256+8+8 Cat-*-main--012.tmp.png montage_012.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--013.tmp.png Cat-cat-cat-main--013.debug.png Cat-cat-cat-main--013.debug.png
convert Cat-cat-cat-main--013.debug.png -bordercolor white -border 8 Cat-cat-cat-main--013.debug.png
montage -geometry 256x256+8+8 Cat-*-main--013.tmp.png montage_013.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--014.tmp.png Cat-cat-cat-main--014.debug.png Cat-cat-cat-main--014.debug.png
convert Cat-cat-cat-main--014.debug.png -bordercolor white -border 8 Cat-cat-cat-main--014.debug.png
montage -geometry 256x256+8+8 Cat-*-main--014.tmp.png montage_014.png
montage -geometry 256x256+8+8 Cat-cat-cat-main--015.tmp.png Cat-cat-cat-main--015.debug.png Cat-cat-cat-main--015.debug.png
convert Cat-cat-cat-main--015.debug.png -bordercolor white -border 8 Cat-cat-cat-main--015.debug.png
montage -geometry 256x256+8+8 Cat-*-main--015.tmp.png montage_015.png
convert Cat-cat-cat-main--*.tmp.png Cat_cat-cat_main.gif
convert Cat-cat-cat-main--*.debug.png Cat_cat-cat_main_debug.gif
montage -geometry 256x256+8+8 Cat-doraemon-main--000.tmp.png Cat-doraemon-main--000.debug.png Cat-doraemon-main--000.debug.png
convert Cat-doraemon-main--000.debug.png -bordercolor white -border 8 Cat-doraemon-main--000.debug.png
montage -geometry 256x256+8+8 Cat-*-main--000.tmp.png montage_000.png
montage -geometry 256x256+8+8 Cat-doraemon-main--001.tmp.png Cat-doraemon-main--001.debug.png Cat-doraemon-main--001.debug.png
convert Cat-doraemon-main--001.debug.png -bordercolor white -border 8 Cat-doraemon-main--001.debug.png
montage -geometry 256x256+8+8 Cat-*-main--001.tmp.png montage_001.png
montage -geometry 256x256+8+8 Cat-doraemon-main--002.tmp.png Cat-doraemon-main--002.debug.png Cat-doraemon-main--002.debug.png
convert Cat-doraemon-main--002.debug.png -bordercolor white -border 8 Cat-doraemon-main--002.debug.png
montage -geometry 256x256+8+8 Cat-*-main--002.tmp.png montage_002.png
montage -geometry 256x256+8+8 Cat-doraemon-main--003.tmp.png Cat-doraemon-main--003.debug.png Cat-doraemon-main--003.debug.png
convert Cat-doraemon-main--003.debug.png -bordercolor white -border 8 Cat-doraemon-main--003.debug.png
montage -geometry 256x256+8+8 Cat-*-main--003.tmp.png montage_003.png
montage -geometry 256x256+8+8 Cat-doraemon-main--004.tmp.png Cat-doraemon-main--004.debug.png Cat-doraemon-main--004.debug.png
convert Cat-doraemon-main--004.debug.png -bordercolor white -border 8 Cat-doraemon-main--004.debug.png
montage -geometry 256x256+8+8 Cat-*-main--004.tmp.png montage_004.png
montage -geometry 256x256+8+8 Cat-doraemon-main--005.tmp.png Cat-doraemon-main--005.debug.png Cat-doraemon-main--005.debug.png
convert Cat-doraemon-main--005.debug.png -bordercolor white -border 8 Cat-doraemon-main--005.debug.png
montage -geometry 256x256+8+8 Cat-*-main--005.tmp.png montage_005.png
montage -geometry 256x256+8+8 Cat-doraemon-main--006.tmp.png Cat-doraemon-main--006.debug.png Cat-doraemon-main--006.debug.png
convert Cat-doraemon-main--006.debug.png -bordercolor white -border 8 Cat-doraemon-main--006.debug.png
montage -geometry 256x256+8+8 Cat-*-main--006.tmp.png montage_006.png
montage -geometry 256x256+8+8 Cat-doraemon-main--007.tmp.png Cat-doraemon-main--007.debug.png Cat-doraemon-main--007.debug.png
convert Cat-doraemon-main--007.debug.png -bordercolor white -border 8 Cat-doraemon-main--007.debug.png
montage -geometry 256x256+8+8 Cat-*-main--007.tmp.png montage_007.png
montage -geometry 256x256+8+8 Cat-doraemon-main--008.tmp.png Cat-doraemon-main--008.debug.png Cat-doraemon-main--008.debug.png
convert Cat-doraemon-main--008.debug.png -bordercolor white -border 8 Cat-doraemon-main--008.debug.png
montage -geometry 256x256+8+8 Cat-*-main--008.tmp.png montage_008.png
montage -geometry 256x256+8+8 Cat-doraemon-main--009.tmp.png Cat-doraemon-main--009.debug.png Cat-doraemon-main--009.debug.png
convert Cat-doraemon-main--009.debug.png -bordercolor white -border 8 Cat-doraemon-main--009.debug.png
montage -geometry 256x256+8+8 Cat-*-main--009.tmp.png montage_009.png
montage -geometry 256x256+8+8 Cat-doraemon-main--010.tmp.png Cat-doraemon-main--010.debug.png Cat-doraemon-main--010.debug.png
convert Cat-doraemon-main--010.debug.png -bordercolor white -border 8 Cat-doraemon-main--010.debug.png
montage -geometry 256x256+8+8 Cat-*-main--010.tmp.png montage_010.png
montage -geometry 256x256+8+8 Cat-doraemon-main--011.tmp.png Cat-doraemon-main--011.debug.png Cat-doraemon-main--011.debug.png
convert Cat-doraemon-main--011.debug.png -bordercolor white -border 8 Cat-doraemon-main--011.debug.png
montage -geometry 256x256+8+8 Cat-*-main--011.tmp.png montage_011.png
montage -geometry 256x256+8+8 Cat-doraemon-main--012.tmp.png Cat-doraemon-main--012.debug.png Cat-doraemon-main--012.debug.png
convert Cat-doraemon-main--012.debug.png -bordercolor white -border 8 Cat-doraemon-main--012.debug.png
montage -geometry 256x256+8+8 Cat-*-main--012.tmp.png montage_012.png
montage -geometry 256x256+8+8 Cat-doraemon-main--013.tmp.png Cat-doraemon-main--013.debug.png Cat-doraemon-main--013.debug.png
convert Cat-doraemon-main--013.debug.png -bordercolor white -border 8 Cat-doraemon-main--013.debug.png
montage -geometry 256x256+8+8 Cat-*-main--013.tmp.png montage_013.png
montage -geometry 256x256+8+8 Cat-doraemon-main--014.tmp.png Cat-doraemon-main--014.debug.png Cat-doraemon-main--014.debug.png
convert Cat-doraemon-main--014.debug.png -bordercolor white -border 8 Cat-doraemon-main--014.debug.png
montage -geometry 256x256+8+8 Cat-*-main--014.tmp.png montage_014.png
montage -geometry 256x256+8+8 Cat-doraemon-main--015.tmp.png Cat-doraemon-main--015.debug.png Cat-doraemon-main--015.debug.png
convert Cat-doraemon-main--015.debug.png -bordercolor white -border 8 Cat-doraemon-main--015.debug.png
montage -geometry 256x256+8+8 Cat-*-main--015.tmp.png montage_015.png
convert Cat-doraemon-main--*.tmp.png Cat_doraemon_main.gif
convert Cat-doraemon-main--*.debug.png Cat_doraemon_main_debug.gif
convert montage_*.png Cat_montage.gif
convert Cat_montage.gif -bordercolor white -border 8 Cat_montage.gif
rm *.tmp.png
rm *.debug.png
rm montage_*.png