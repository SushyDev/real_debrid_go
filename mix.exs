defmodule RealDebrid.MixProject do
  use Mix.Project

  def project do
    [
      app: :real_debrid,
      version: "0.1.0",
      elixir: "~> 1.14",
      start_permanent: Mix.env() == :prod,
      deps: deps(),
      description: "Elixir client for the Real-Debrid API",
      package: package(),
      source_url: "https://github.com/sushydev/real_debrid_go"
    ]
  end

  def application do
    [
      extra_applications: [:logger]
    ]
  end

  defp deps do
    [
      {:req, "~> 0.4"},
      {:jason, "~> 1.4"}
    ]
  end

  defp package do
    [
      licenses: ["MIT"],
      links: %{"GitHub" => "https://github.com/sushydev/real_debrid_go"}
    ]
  end
end
