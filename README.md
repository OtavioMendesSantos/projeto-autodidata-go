# Projeto Autodidata Go

Estarei utilizando o curso [Aprenda Go 🇧🇷](https://youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&si=HW2PTfFI8oqfarj1);

[Repositório](https://github.com/vkorbes/aprendago) com os materiais do curso;

# Introdução

## Instalação do Go no Linux

1. **Baixe o Go oficial**

   Acesse o site: https://go.dev/dl/

2. **Remova instalações antigas (se houver)**

   ```bash
   sudo rm -rf /usr/local/go
   ```

3. **Extraia o Go em `/usr/local`**

   ```bash
   sudo tar -C /usr/local -xzf go1.25.6.linux-amd64.tar.gz
   ```

4. **Adicione o Go ao `PATH`**

   Edite o arquivo de configuração do shell:

   ```bash
   nano ~/.bashrc
   ```

   Adicione no final:

   ```bash
   export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
   ```

5. **Recarregue o shell**

   ```bash
   source ~/.bashrc
   ```

6. **Verifique a instalação**

   ```bash
   go version
   ```

   Saída esperada:

   ```text
   go version go1.25.6 linux/amd64
   ```

Pronto. O Go está instalado e pronto para uso.
