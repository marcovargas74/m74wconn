#!/bin/sh

DOWNLOAD_SERVER=http://marcovargas.com.br/Downloads

unzipPush(){
	cd /usr/bin
	unzip whatsapp.zip
	rm -f whatsapp.zip
}

#Instala Push usado em notificacao dos aplicativos IoS
#Foi feito desta maneira provisoriamente 
#porque nao tem espaço no Firmware  
installWhatsapp(){
  if [ -e /usr/bin/whatsapp ]; then
    #echo "boot_unniti at `date` " >> /tmp/boot_ok
    logger -p local0.info "instalado o WhatsBot"
    rm -rf /usr/bin/whatsapp
    #return
  fi

  #Muda Pasta para baixar arquivos
  cd /tmp/
  #baixa arquivo
  wget $DOWNLOAD_SERVER/whatsapp.zip
  if [ -e /tmp/whatsapp.zip ]; then
	  mv /tmp/whatsapp.zip /usr/bin/whatsapp.zip
	  cd /usr/bin 
    unzipPush
    chmod 755 /usr/bin/whatsapp
    #Copia os certificados
    #mv /usr/bin/serverAPN* /etc/ssl/private/
    #chmod 600 /etc/ssl/private/serverAPN*
	  logger -p local0.info "Sucesso ao instalar o WHATSAPP_BOT"
  else
	  logger -p local0.err "Erro ao buscar o SEND_PUSH em $DOWNLOAD_SERVER"
  fi
 
}


echo "last_update whatsapp at `date` " >> /data/logs/whatsapp.log
#echo "last_update whatsapp at `date` " > whatsapp.log

installWhatsapp 
sleep 2
sh /etc/init.d/S74whatsapp stop


