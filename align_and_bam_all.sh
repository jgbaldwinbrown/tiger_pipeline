#!/bin/bash
set -e

GENOME=

if [ ! -f ${GENOME}.indexed ] ; then
	bwa index ${GENOME}
	touch ${GENOME}.indexed
fi

# inputs.txt contains the names of the input fastq files -- only the forward
# and reverse reads (R1 and R3). One sample per line, with a tab between the forward and
# reverse names.
cat inputs.txt | while read line; do
	F=`echo $line | cut -d '	' -f 1`
	R=`echo $line | cut -d '	' -f 2`
	bwa mem -t 8 ${GENOME} $F $R | samtools -Sb > ${F}.bam
done
