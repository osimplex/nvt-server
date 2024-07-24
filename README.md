# Network Virtual Terminal Server

Um protótipo de servidor de interfaces TUI (*terminal user interface*) baseada em registradores e chamada de procedure para terminais remotos. Objetivando flexibilidade e extensibilidade, simplificando a substituição de componentes na arquitetura de forma a possibilitar que um *worker* ou *handler* seja associado com qualquer *server* disponível.

## Uso

Basta executar `go run nvt-server` para executar um modo de demonstração mínimo, cujo arquivo de configuração está em `etc/example.toml`.

## Licença

O *nvt-server* está disponível sob a licença AGPL3.

## Histórico

Este protótipo foi criado para uma solução com dispositivos embarcados de terminal remoto, possibilitando acesso a um sistema informatizado sem a necessidade de um *smartphone* ou outro dispositivo mais complexo. A comunicação com os dispositivos embarcados realizou-se com o protocolo Telnet somente para fins de agilizar a prototipação.

Neste protótipo provou-se o conceito de uma estrutura de interfaces TUI baseada em comandos simples para a execução de tarefas através de algum sistema de processos remotos. Com *procedures* e tabelas num SGBD Oracle está estruturada a demonstração, mas pode-se utilizar qualquer outra tecnologia como gRPC ou APIs REST, bastando implementar o devido *worker* e realizar a configuração.
