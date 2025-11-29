defmodule RealDebrid.Api.TorrentsExtras do
  @moduledoc """
  Additional torrent-related API functions.
  """

  alias RealDebrid.Client

  @type active_count :: %{
          nb: integer(),
          limit: integer()
        }

  @type available_host :: %{
          host: String.t(),
          max_file_size: integer()
        }

  @doc """
  Get the count of active torrents and the limit.

  ## Returns

    - `{:ok, %{nb: count, limit: limit}}` on success
    - `{:error, reason}` on failure
  """
  @spec get_active_count(Client.t()) :: {:ok, active_count()} | {:error, term()}
  def get_active_count(%Client{} = client) do
    case Client.get(client, "/torrents/activeCount") do
      {:ok, body, _headers} ->
        {:ok,
         %{
           nb: body["nb"],
           limit: body["limit"]
         }}

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Get list of available hosts for torrents.

  ## Returns

    - `{:ok, hosts}` - List of available hosts
    - `{:error, reason}` on failure
  """
  @spec get_available_hosts(Client.t()) :: {:ok, [available_host()]} | {:error, term()}
  def get_available_hosts(%Client{} = client) do
    case Client.get(client, "/torrents/availableHosts") do
      {:ok, body, _headers} ->
        hosts =
          Enum.map(body, fn host ->
            %{
              host: host["host"],
              max_file_size: host["max_file_size"]
            }
          end)

        {:ok, hosts}

      {:error, reason} ->
        {:error, reason}
    end
  end
end
