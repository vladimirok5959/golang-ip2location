#!/bin/bash

TOKEN="$1"
FILE="DB3LITEBIN"
URL="https://www.ip2location.com/download/?token=${TOKEN}&file=${FILE}"

FILE_ZIP_SOURCE="/tmp/db.zip"
FILE_BIN_SOURCE="/tmp/IP2LOCATION-LITE-DB3.BIN"
FILE_BIN_TARGET="$2"

if [ "${TOKEN}" = "" ]; then
	echo "Please, set token"
	exit 1
fi

# Download
echo "Download ip2location file..."
curl -s -o "${FILE_ZIP_SOURCE}" "${URL}"

# Check
if [ ! -f "${FILE_ZIP_SOURCE}" ]; then
	echo "ZIP file is not exists"
	exit 1
fi

SIZE=$(stat --printf="%s" "${FILE_ZIP_SOURCE}")
if [ ${SIZE} -lt 1024 ]; then
	echo "ZIP file is small: ${SIZE}"
	cat "${FILE_ZIP_SOURCE}"
	echo "" && rm "${FILE_ZIP_SOURCE}"
	exit 1
fi

# Unpack
unzip -p "${FILE_ZIP_SOURCE}" "IP2LOCATION-LITE-DB3.BIN" > "${FILE_BIN_SOURCE}"
rm "${FILE_ZIP_SOURCE}"

# Check
if [ ! -f "${FILE_BIN_SOURCE}" ]; then
	echo "BIN file is not exists"
	exit 1
fi

SIZE=$(stat --printf="%s" "${FILE_BIN_SOURCE}")
if [ ${SIZE} -lt 1024 ]; then
	echo "BIN file is small: ${SIZE}"
	rm "${FILE_BIN_SOURCE}"
	exit 1
fi

# Replace
mv "${FILE_BIN_SOURCE}" "${FILE_BIN_TARGET}"
