defmodule RealDebrid do
  @moduledoc """
  A client library for interacting with the Real-Debrid API.

  ## Usage

      # Create a client with your API token
      client = RealDebrid.Client.new("your_api_token")

      # Get user information
      {:ok, user} = RealDebrid.Api.User.get(client)

      # Get torrents list
      {:ok, %{torrents: torrents}} = RealDebrid.Api.Torrents.get(client, limit: 100, page: 1)

      # Add a magnet link
      {:ok, result} = RealDebrid.Api.AddMagnet.add(client, "magnet:?xt=...")
  """
end
