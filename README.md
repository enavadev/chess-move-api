Chess Move API - GO Lang
========================

Esse projeto é um exemplo de uma API (Back-end) que lista os próximos moviemtnos do Cavalo a partir de uma determinada posição (em notção agébrica)
		
	- Skills:
	
    #Nesse projeto foram utilizados as seguintes ferramentas:

		* Utilizada IDE Visual Studio Code;
		* Go lang versão 1.17.5 com go mod;
		* Banco de dados utilizado é o MongoDB (para salvar os logs);
        * Framework GIN GONIC para prover HTTP Server. Eu acho uma dos melhores packages para esse serviço, tem uma  performance excelente;
        * Docker e Docker-Compose para subir os container da aplicação e do Mongo;

    # INSTALAÇÃO DA API:

        - Clone o projeto na sua máquina;
        - Já com o docker e docker-compose instalados (é necessáraio que estejam devidamente instalados na sua máquina), é só rodar o camando abaixo no shell do seu S.O.:
            > docker-compose up -d
        - Espere o processo concluir e tudo ocorrendo 100%, a API iniciará na porta 4000;
        - No seu navegador informe como endereço a seguinte url:
            > http://localhost:4000/api/v1/knight/<posição>
                - Onde: <posição> -> corresponde a posição do tabuleiro de xadrez em notação algébrica por exemplo:
                   .../v1/knight/D4
