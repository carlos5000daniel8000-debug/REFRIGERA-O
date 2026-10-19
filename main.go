package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    html := `<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Refrigeração Palmeira | Soluções em Climatização</title>
    <style>
        :root {
            --primary: #0b5d7a;
            --primary-dark: #083d50;
            --secondary: #28b8d9;
            --accent: #f3b93d;
            --bg: #edf7fb;
            --surface: #ffffff;
            --text: #16313f;
            --muted: #58707d;
            --line: #d9edf4;
        }
        * { box-sizing: border-box; margin: 0; padding: 0; }
        html { scroll-behavior: smooth; }
        body {
            font-family: Arial, sans-serif;
            background: linear-gradient(180deg, #f6fbfd 0%, #eef8fb 100%);
            color: var(--text);
            line-height: 1.6;
        }
        a { text-decoration: none; }
        .container { width: min(1120px, calc(100% - 32px)); margin: 0 auto; }
        .topbar {
            background: rgba(11, 93, 122, 0.96);
            color: white;
            position: sticky;
            top: 0;
            z-index: 10;
            backdrop-filter: blur(8px);
        }
        .topbar-inner {
            display: flex;
            align-items: center;
            justify-content: space-between;
            padding: 18px 0;
            gap: 20px;
        }
        .brand {
            font-size: 1.25rem;
            font-weight: bold;
            letter-spacing: 0.02em;
        }
        .nav {
            display: flex;
            align-items: center;
            gap: 22px;
            flex-wrap: wrap;
        }
        .nav a {
            color: rgba(255,255,255,0.88);
            font-size: 0.96rem;
        }
        .nav a:hover { color: white; }
        .btn {
            display: inline-flex;
            align-items: center;
            justify-content: center;
            border-radius: 999px;
            padding: 14px 24px;
            font-weight: bold;
            transition: 0.2s ease;
        }
        .btn-primary {
            background: var(--accent);
            color: var(--primary-dark);
            box-shadow: 0 12px 26px rgba(243, 185, 61, 0.35);
        }
        .btn-primary:hover { transform: translateY(-2px); }
        .btn-whatsapp {
            background: #25d366;
            color: white;
            border: none;
            box-shadow: 0 12px 26px rgba(37, 211, 102, 0.25);
        }
        .btn-secondary {
            background: white;
            color: var(--primary);
            border: 1px solid rgba(11,93,122,0.15);
        }
        .hero {
            padding: 72px 0 40px;
            background: linear-gradient(135deg, var(--primary-dark) 0%, var(--primary) 48%, var(--secondary) 100%);
            color: white;
        }
        .hero-inner {
            display: grid;
            grid-template-columns: 1.2fr 0.8fr;
            gap: 34px;
            align-items: center;
        }
        .eyebrow {
            display: inline-block;
            background: rgba(255,255,255,0.12);
            border: 1px solid rgba(255,255,255,0.18);
            padding: 8px 12px;
            border-radius: 999px;
            font-size: 0.8rem;
            letter-spacing: 0.08em;
            text-transform: uppercase;
            margin-bottom: 18px;
        }
        h1 {
            font-size: clamp(2.3rem, 5vw, 4rem);
            line-height: 1.1;
            margin-bottom: 18px;
        }
        .hero p {
            max-width: 630px;
            color: rgba(255,255,255,0.86);
            font-size: 1.08rem;
            margin-bottom: 28px;
        }
        .hero-actions {
            display: flex;
            gap: 16px;
            align-items: center;
            flex-wrap: wrap;
        }
        .hero-list {
            display: flex;
            gap: 18px;
            flex-wrap: wrap;
            margin-top: 24px;
        }
        .hero-list span {
            background: rgba(255,255,255,0.08);
            border: 1px solid rgba(255,255,255,0.12);
            padding: 10px 14px;
            border-radius: 999px;
            font-size: 0.9rem;
        }
        .card-visual {
            background: rgba(255,255,255,0.08);
            border: 1px solid rgba(255,255,255,0.16);
            border-radius: 28px;
            padding: 20px;
            box-shadow: 0 30px 60px rgba(2, 24, 40, 0.25);
        }
        .card-visual img {
            display: block;
            width: 100%;
            border-radius: 20px;
            height: 360px;
            object-fit: cover;
        }
        .stats {
            display: grid;
            grid-template-columns: repeat(3, minmax(0, 1fr));
            gap: 18px;
            margin-top: -28px;
            position: relative;
            z-index: 2;
        }
        .stat {
            background: white;
            border: 1px solid var(--line);
            border-radius: 20px;
            padding: 26px 18px;
            text-align: center;
            box-shadow: 0 18px 36px rgba(11, 93, 122, 0.08);
        }
        .stat strong {
            display: block;
            font-size: 2rem;
            color: var(--primary);
        }
        .stat span {
            color: var(--muted);
            font-size: 0.96rem;
        }
        section {
            padding: 88px 0;
        }
        .section-title {
            text-align: center;
            margin-bottom: 34px;
        }
        .section-title h2 {
            font-size: clamp(2rem, 4vw, 2.7rem);
            margin-bottom: 12px;
        }
        .section-title p {
            max-width: 700px;
            margin: 0 auto;
            color: var(--muted);
            font-size: 1.05rem;
        }
        .services-grid,
        .benefits-grid,
        .process-grid {
            display: grid;
            grid-template-columns: repeat(4, minmax(0, 1fr));
            gap: 22px;
        }
        .service,
        .benefit,
        .process {
            background: var(--surface);
            border: 1px solid var(--line);
            border-radius: 22px;
            padding: 26px 20px;
            box-shadow: 0 14px 30px rgba(11, 93, 122, 0.05);
        }
        .service .icon,
        .benefit .icon,
        .process .icon {
            width: 54px;
            height: 54px;
            border-radius: 16px;
            display: grid;
            place-items: center;
            font-size: 1.5rem;
            background: rgba(40, 184, 217, 0.1);
            color: var(--primary);
            margin-bottom: 18px;
        }
        .service h3,
        .benefit h3,
        .process h3 {
            margin-bottom: 10px;
            font-size: 1.3rem;
        }
        .service p,
        .benefit p,
        .process p,
        .contact-card p {
            color: var(--muted);
        }
        .feature-panel {
            background: white;
            border: 1px solid var(--line);
            border-radius: 30px;
            padding: 28px;
            box-shadow: 0 18px 46px rgba(11, 93, 122, 0.08);
        }
        .feature-layout {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 30px;
            align-items: center;
        }
        .feature-list {
            list-style: none;
            display: grid;
            gap: 16px;
        }
        .feature-list li {
            display: flex;
            align-items: flex-start;
            gap: 12px;
            color: var(--muted);
            font-size: 1rem;
        }
        .check {
            width: 22px;
            height: 22px;
            border-radius: 50%;
            background: rgba(40, 184, 217, 0.15);
            display: grid;
            place-items: center;
            color: var(--primary);
            font-weight: bold;
            flex-shrink: 0;
            margin-top: 2px;
        }
        .contact-wrap {
            display: grid;
            grid-template-columns: 1.1fr 0.9fr;
            gap: 26px;
        }
        .contact-card {
            background: white;
            border: 1px solid var(--line);
            border-radius: 24px;
            padding: 28px;
            box-shadow: 0 16px 36px rgba(11, 93, 122, 0.06);
        }
        .contact-item {
            display: flex;
            flex-wrap: wrap;
            gap: 12px;
            padding: 14px 0;
            border-bottom: 1px solid var(--line);
        }
        .contact-item:last-child { border-bottom: none; }
        .contact-item strong {
            width: 100px;
            color: var(--text);
        }
        .qr-box {
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 18px 0 6px;
        }
        .qr-box a {
            display: inline-block;
            background: white;
            border: 1px solid var(--line);
            border-radius: 18px;
            padding: 12px;
            box-shadow: 0 10px 28px rgba(11, 93, 122, 0.08);
        }
        .qr-box img {
            display: block;
            width: 180px;
            height: 180px;
            border-radius: 12px;
        }
        form {
            display: grid;
            gap: 14px;
            background: white;
            border: 1px solid var(--line);
            border-radius: 24px;
            padding: 28px;
            box-shadow: 0 16px 36px rgba(11, 93, 122, 0.06);
        }
        form input,
        form textarea {
            width: 100%;
            padding: 14px 16px;
            border-radius: 12px;
            border: 1px solid var(--line);
            background: #f9fcfd;
            font: inherit;
            color: var(--text);
        }
        form textarea { min-height: 120px; resize: vertical; }
        .hidden { display: none; }
        .footer {
            background: var(--primary-dark);
            color: rgba(255,255,255,0.8);
            text-align: center;
            padding: 22px 0;
            font-size: 0.95rem;
        }
        @media (max-width: 900px) {
            .hero-inner,
            .feature-layout,
            .contact-wrap,
            .services-grid,
            .benefits-grid,
            .process-grid {
                grid-template-columns: 1fr 1fr;
            }
            .services-grid, .benefits-grid, .process-grid { grid-template-columns: repeat(2, minmax(0, 1fr)); }
            .hero-inner, .feature-layout, .contact-wrap { grid-template-columns: 1fr; }
        }
        @media (max-width: 640px) {
            .topbar-inner {
                flex-direction: column;
                align-items: flex-start;
            }
            .nav { gap: 12px; }
            .stats,
            .services-grid,
            .benefits-grid,
            .process-grid {
                grid-template-columns: 1fr;
            }
            .hero { padding-top: 52px; }
            .hero-actions {
                flex-direction: column;
                align-items: stretch;
            }
            .btn { width: 100%; }
        }
    </style>
</head>
<body>
    <header class="topbar">
        <div class="container topbar-inner">
            <div class="brand">Refrigeração Palmeira</div>
            <nav class="nav" aria-label="Menu principal">
                <a href="#inicio">Início</a>
                <a href="#servicos">Serviços</a>
                <a href="#sobre">Sobre</a>
                <a href="#contato">Contato</a>
            </nav>
        </div>
    </header>

    <main>
        <section class="hero" id="inicio">
            <div class="container hero-inner">
                <div>
                    <span class="eyebrow">Conforto e eficiência</span>
                    <h1>Climatização e refrigeração que você pode confiar.</h1>
                    <p>Serviços profissionais em refrigeração residencial e comercial. Instalação, manutenção, conserto e limpeza de aparelhos de ar-condicionado, câmaras frias e geladeiras com atendimento rápido e garantia de qualidade.</p>
                    <div class="hero-actions">
                        <a class="btn btn-primary" href="https://wa.me/5583988354958?text=Ol%C3%A1%2C%20quero%20solicitar%20um%20or%C3%A7amento." target="_blank" rel="noopener noreferrer">Solicitar orçamento</a>
                        <a class="btn btn-secondary" href="#servicos">Conhecer serviços</a>
                    </div>
                    <div class="hero-list">
                        <span>Instalação</span>
                        <span>Manutenção</span>
                        <span>Reparo</span>
                        <span>Limpeza</span>
                    </div>
                </div>
                <div class="card-visual">
                    <img src="/fundo.png.jpg" alt="Refrigeração Palmeira" />
                </div>
            </div>
        </section>

        <div class="container stats">
            <div class="stat">
                <strong>+10 anos</strong>
                <span>de experiência</span>
            </div>
            <div class="stat">
                <strong>1.200+</strong>
                <span>atendimentos realizados</span>
            </div>
            <div class="stat">
                <strong>24h</strong>
                <span>para suporte rápido</span>
            </div>
        </div>

        <section id="servicos">
            <div class="container">
                <div class="section-title">
                    <h2>Nossos serviços</h2>
                    <p>Atendimento completo para manter seu ambiente climatizado e seus equipamentos funcionando com eficiência o ano todo.</p>
                </div>
                <div class="services-grid">
                    <article class="service">
                        <div class="icon">❄️</div>
                        <h3>Instalação</h3>
                        <p>Instalação de aparelhos split, janela e cassete com mão de obra especializada e segurança.</p>
                    </article>
                    <article class="service">
                        <div class="icon">🛠️</div>
                        <h3>Manutenção</h3>
                        <p>Manutenção preventiva e corretiva para garantir desempenho ideal e baixa de energia.</p>
                    </article>
                    <article class="service">
                        <div class="icon">🔧</div>
                        <h3>Reparo</h3>
                        <p>Diagnóstico rápido para vazamentos, peças, revisões elétricas e correções emergenciais.</p>
                    </article>
                    <article class="service">
                        <div class="icon">🧼</div>
                        <h3>Limpeza Técnica</h3>
                        <p>Limpeza completa interna e externa com remoção de sujeira, fungos e odores.</p>
                    </article>
                </div>
            </div>
        </section>

        <section id="sobre">
            <div class="container feature-panel">
                <div class="section-title">
                    <h2>Por que escolher a Refrigeração Palmeira?</h2>
                    <p>Experiência, confiança e atendimento local para manter seu conforto em perfeito estado.</p>
                </div>
                <div class="feature-layout">
                    <div>
                        <ul class="feature-list">
                            <li><span class="check">✓</span><span><strong>Equipe especializada:</strong> técnicos treinados e prontos para resolver qualquer necessidade com segurança.</span></li>
                            <li><span class="check">✓</span><span><strong>Atendimento ágil:</strong> respostas rápidas para emergências e suporte proativo para evitar maiores problemas.</span></li>
                            <li><span class="check">✓</span><span><strong>Qualidade garantida:</strong> peças adequadas, orçamento claro e serviço finalizado com foco na satisfação do cliente.</span></li>
                            <li><span class="check">✓</span><span><strong>Preço justo:</strong> soluções práticas para residências, comércios e ambientes profissionais.</span></li>
                        </ul>
                    </div>
                    <div class="benefits-grid">
                        <div class="benefit">
                            <div class="icon">🏆</div>
                            <h3>Confiabilidade</h3>
                            <p>Trabalho bem executado e atenção aos detalhes.</p>
                        </div>
                        <div class="benefit">
                            <div class="icon">⚡</div>
                            <h3>Rapidez</h3>
                            <p>Agilidade para resolver falhas e minimizar impactos.</p>
                        </div>
                        <div class="benefit">
                            <div class="icon">🌡️</div>
                            <h3>Eficiência</h3>
                            <p>Ambientes mais confortáveis e menor consumo energético.</p>
                        </div>
                        <div class="benefit">
                            <div class="icon">🤝</div>
                            <h3>Relacionamento</h3>
                            <p>Atendimento humano, transparente e comprometido.</p>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <section>
            <div class="container">
                <div class="section-title">
                    <h2>Como trabalhamos</h2>
                    <p>Processo simples, rápido e pensado para o melhor resultado.</p>
                </div>
                <div class="process-grid">
                    <div class="process">
                        <div class="icon">1</div>
                        <h3>Contato</h3>
                        <p>Você entra em contato e descreve a necessidade ou a ocorrência.</p>
                    </div>
                    <div class="process">
                        <div class="icon">2</div>
                        <h3>Diagnóstico</h3>
                        <p>Realizamos análise do sistema e indicamos a melhor solução.</p>
                    </div>
                    <div class="process">
                        <div class="icon">3</div>
                        <h3>Execução</h3>
                        <p>Fazemos a instalação, manutenção ou reparo com atenção e cuidado.</p>
                    </div>
                    <div class="process">
                        <div class="icon">4</div>
                        <h3>Garantia</h3>
                        <p>Entrega do serviço com suporte e compromisso com seu conforto.</p>
                    </div>
                </div>
            </div>
        </section>

        <section id="contato">
            <div class="container contact-wrap">
                <div class="contact-card">
                    <div class="section-title" style="text-align:left; margin-bottom:18px;">
                        <h2>Fale com a gente</h2>
                    </div>
                    <p>Pronto para atender sua necessidade de refrigeração e climatização com rapidez, qualidade e profissionalismo.</p>
                    <div class="contact-item"><strong>Telefone:</strong><span>(83) 9 8835-4958</span></div>
                    <div class="contact-item"><strong>Telefone:</strong><span>(83) 9 8734-8115</span></div>
                    <div class="contact-item"><strong>Email:</strong><span>contato@refrigeracaopalmeira.com.br</span></div>
                    <div class="contact-item"><strong>Atendimento:</strong><span>Seg a Sab, 08h às 18h</span></div>
                    <div class="contact-item"><strong>Local:</strong><span>Palmeira, PB</span></div>
                    <div class="qr-box">
                        <a href="https://wa.me/5583988354958?text=Ol%C3%A1%2C%20vim%20pelo%20site%20e%20quero%20um%20or%C3%A7amento." target="_blank" rel="noopener noreferrer" aria-label="Abrir WhatsApp">
                            <img src="https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=https://wa.me/5583988354958?text=Ol%C3%A1%2C%20vim%20pelo%20site%20e%20quero%20um%20or%C3%A7amento." alt="QR Code WhatsApp" />
                        </a>
                    </div>
                </div>

                <form id="orcamentoForm">
                    <h3 style="font-size:1.5rem;">Solicitar orçamento</h3>
                    <input type="text" id="nome" placeholder="Seu nome" required />
                    <input type="text" id="telefone" placeholder="Seu telefone" required />
                    <input type="email" id="email" placeholder="Seu e-mail" required />
                    <textarea id="mensagem" placeholder="Descreva o que você precisa" required></textarea>
                    <button type="submit" class="btn btn-whatsapp" style="cursor:pointer;">Enviar pedido</button>
                </form>

                <script>
                    document.getElementById('orcamentoForm').addEventListener('submit', function (event) {
                        event.preventDefault();

                        const nome = document.getElementById('nome').value.trim();
                        const telefone = document.getElementById('telefone').value.trim();
                        const email = document.getElementById('email').value.trim();
                        const mensagem = document.getElementById('mensagem').value.trim();

                        const texto = encodeURIComponent(
                            'Olá, quero solicitar um orçamento.\n\n' +
                            'Nome: ' + nome + '\n' +
                            'Telefone: ' + telefone + '\n' +
                            'E-mail: ' + email + '\n' +
                            'Mensagem: ' + mensagem
                        );

                        window.open('https://wa.me/5583988354958?text=' + texto, '_blank');
                    });
                </script>
            </div>
        </section>
    </main>

    <footer class="footer">
        Refrigeração Palmeira © 2026 - Soluções em refrigeração e climatização.
    </footer>
</body>
</html>`

    fmt.Fprint(w, html)
}

func main() {
    host := os.Getenv("HOST")
    if host == "" {
        host = "0.0.0.0"
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    addr := net.JoinHostPort(host, port)

    http.Handle("/fundo.png.jpg", http.FileServer(http.Dir(".")))
    http.HandleFunc("/", homeHandler)

    fmt.Printf("Servidor rodando em http://%s\n", addr)
    fmt.Printf("Também pode ser acessado via localhost:%s\n", port)
    log.Fatal(http.ListenAndServe(addr, nil))
}
