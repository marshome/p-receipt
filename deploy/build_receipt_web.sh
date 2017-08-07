#!/usr/bin/env bash

rm ../_output/receipt-web.zip
cd ../frontend/receipt-web
rm -rf build
npm run build
zip -r ../../_output/receipt-web ./build