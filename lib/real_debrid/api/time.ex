defmodule RealDebrid.Api.Time do
  @moduledoc """
  Server time API functions.
  """

  alias RealDebrid.Client

  @doc """
  Get the server time as a string.

  ## Returns

    - `{:ok, time}` - Server time string
    - `{:error, reason}` on failure
  """
  @spec get(Client.t()) :: {:ok, String.t()} | {:error, term()}
  def get(%Client{} = client) do
    case Client.get(client, "/time") do
      {:ok, body, _headers} when is_binary(body) ->
        {:ok, body}

      {:ok, body, _headers} ->
        {:ok, to_string(body)}

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Get the server time in ISO 8601 format.

  ## Returns

    - `{:ok, time}` - Server time in ISO 8601 format
    - `{:error, reason}` on failure
  """
  @spec get_iso(Client.t()) :: {:ok, String.t()} | {:error, term()}
  def get_iso(%Client{} = client) do
    case Client.get(client, "/time/iso") do
      {:ok, body, _headers} when is_binary(body) ->
        {:ok, body}

      {:ok, body, _headers} ->
        {:ok, to_string(body)}

      {:error, reason} ->
        {:error, reason}
    end
  end
end
