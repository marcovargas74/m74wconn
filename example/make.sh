#!/bin/sh
#Script usado para Compila aplicação para ARM
#
#
APP_NAME=whatsApp
DIR_WORK_APP=~/projetos-go/src/github.com//marcovargas74/m74wconn/example
#DIR_FIRMWARE=~/UNNITI_GIT/firmware/packetsAPP_bin/release_voip
IP_UNNITI=172.31.0.180


echo "--------------------------INICIO------------------------------------"
echo "1. INICIANDO a Compilacao do app!!!!"

cd $DIR_WORK_APP
env GOOS=linux GOARCH=arm go build -o $APP_NAME

if [ -e $APP_NAME ]; then
  echo "2. OK APP Compilado!!!!"
else
  echo "2. NOK Problema ao compilar APP!!!!"
  exit
fi	

#sleep 1
#echo "3. Zipa aplicacao para enviar para gerador de firmware!!!"
#zip $APP_NAME.zip $APP_NAME

ls -l -h 
#copia a aplicacao para pasta do gerador de firmware
#diretorio usado para gerar as Versoes de Firmware da Versao desenvolvimento
echo ""
echo ""
#cp -f $DIR_WORK_APP/$APP_NAME.zip $DIR_FIRMWARE/$APP_NAME.zip
#echo "4.Copia o arquivo binario $APP_NAME para a pasta do Firmware!!!"
#echo "--------------------------INICIO------------------------------------"
#cd $DIR_WORK_APP
#se Nao existir apaga o antigo 
#if [ ! -e fpga.ko ]; then
# echo "ARQUIVO DO Drive FPGA nao Existe!!! "
#fi	

#TESTE copia firmware para testar na Unniti
echo "Copia o arquivo binario para UNNITI $IP_UNNITI!!!!"
#ls -l

#scp -P 16022 $DIR_WORK_APP/$APP_NAME.zip  dev@$IP_UNNITI:/tmp
scp -P 16022 $DIR_WORK_APP/$APP_NAME  dev@$IP_UNNITI:/tmp

echo "Arquivo copiado para para UNNITI $IP_UNNITI!!!"
echo "--------------------------FIM------------------------------------"
echo ""

cd ..


