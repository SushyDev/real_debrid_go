defmodule RealDebrid.Api.InstantAvailability do
  @moduledoc """
  Check instant availability of a torrent hash.
  """

  alias RealDebrid.Client

  @type file :: %{
          filename: String.t(),
          filesize: integer()
        }

  @type instant_availability :: %{
          hash: String.t(),
          files: map()
        }

  @doc """
  Get instant availability for a torrent hash.

  ## Parameters

    - `client` - The Real-Debrid client
    - `hash` - The torrent hash to check

  ## Returns

    - `{:ok, %{hash: hash, files: files}}` on success
    - `{:error, reason}` on failure
  """
  @spec get(Client.t(), String.t()) :: {:ok, instant_availability()} | {:error, term()}
  def get(%Client{} = client, hash) do
    case Client.get(client, "/torrents/instantAvailability/#{hash}") do
      {:ok, body, _headers} ->
        {:ok,
         %{
           hash: hash,
           files: body
         }}

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Extract files from an instant availability result.
  """
  @spec get_files(instant_availability()) :: map()
  def get_files(%{files: files}), do: files

  @doc """
  Extract hash from an instant availability result.
  """
  @spec get_hash(instant_availability()) :: String.t()
  def get_hash(%{hash: hash}), do: hash
end
