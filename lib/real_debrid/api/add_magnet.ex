defmodule RealDebrid.Api.AddMagnet do
  @moduledoc """
  Add a magnet link to Real-Debrid.
  """

  alias RealDebrid.Client

  @type response :: %{
          id: String.t(),
          uri: String.t()
        }

  @doc """
  Add a magnet link to the user's torrents.

  ## Parameters

    - `client` - The Real-Debrid client
    - `magnet` - The magnet link to add
    - `opts` - Optional keyword list:
      - `:host` - The host to use (optional)

  ## Returns

    - `{:ok, %{id: id, uri: uri}}` on success
    - `{:error, reason}` on failure
  """
  @spec add(Client.t(), String.t(), keyword()) :: {:ok, response()} | {:error, term()}
  def add(%Client{} = client, magnet, opts \\ []) do
    form = %{"magnet" => magnet}

    form =
      case Keyword.get(opts, :host) do
        nil -> form
        host -> Map.put(form, "host", host)
      end

    case Client.post(client, "/torrents/addMagnet", form: form, expected_status: 201) do
      {:ok, body} ->
        {:ok,
         %{
           id: body["id"],
           uri: body["uri"]
         }}

      {:error, reason} ->
        {:error, reason}
    end
  end
end
