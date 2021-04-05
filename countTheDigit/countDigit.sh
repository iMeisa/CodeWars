#!/bin/bash
nbDig(){
    i=0
    declare -a N_DIGITS
    while [ $i -le $1 ]
    do
        isq="$(( i*i ))"
        N_DIGITS+=($isq)
        i=$(( i+1 ))
    done

    count=0
    for var in ${N_DIGITS[*]}
    do
        for num in $(seq 1 ${#var})
        do
            echo $num
            case $num in
                *$2*)
                    count=$(( count+1 ))
            esac
        done
    done
    echo $count
}
nbDig 575 0