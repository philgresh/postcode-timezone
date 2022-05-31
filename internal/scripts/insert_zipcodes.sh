#!/bin/bash
set -e

INITIAL_WORKING_DIRECTORY=$(pwd)
cd "$(dirname "$0")"
INPUT_CSV="../data/zipcodes.us.csv"
OUTPUT=""

function get_state_id() {
  zipcode="$1"
  r=$(sqlite3 ../data/db.sqlite3 "SELECT state_id FROM zipcodes WHERE code = ${zipcode}")
  echo $r
}

function get_county_id() {
  state_id="$1"
  province_code="$2"
  r=$(sqlite3 ../data/db.sqlite3 "SELECT counties.id FROM counties WHERE counties.state_id = \"${state_id}\" AND abbr = \"${province_code}\"")
  echo $r
}

function update_zipcode() {
  zipcode="$1"
  county_id="$2"
  r=$(sqlite3 ../data/db.sqlite3 "UPDATE zipcodes SET county_id = \"${county_id}\" WHERE zipcodes.code = \"${zipcode}\"")
  echo $r
}

while IFS= read -r line
do
  IFS=','
  read -a strarr <<< "$line"
  zipcode="${strarr[1]}"
  province_code="${strarr[6]}"
  state_id="$(get_state_id ${zipcode})"
  county_id="$(get_county_id ${state_id} ${province_code})"

  update_zipcode $zipcode $county_id
  echo "Updating zipcode '$zipcode' with countyID '$county_id' and stateID '$state_id'"
    
  IFS=
done < "$INPUT_CSV"

echo "Done!"
