defmodule RealDebrid.Api.Delete do
  @moduledoc """
  Delete a torrent from Real-Debrid.
  """

  alias RealDebrid.Client

  @doc """
  Delete a torrent by its ID.

  ## Parameters

    - `client` - The Real-Debrid client
    - `id` - The torrent ID to delete

  ## Returns

    - `:ok` on success (204 response)
    - `{:error, reason}` on failure
  """
  @spec delete(Client.t(), String.t()) :: :ok | {:error, term()}
  def delete(%Client{} = client, id) do
    Client.delete(client, "/torrents/delete/#{id}")
  end
end
