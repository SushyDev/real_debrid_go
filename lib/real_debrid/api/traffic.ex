defmodule RealDebrid.Api.Traffic do
  @moduledoc """
  Traffic information API functions.
  """

  alias RealDebrid.Client

  @type traffic_entry :: %{
          left: integer(),
          bytes: integer(),
          links: integer(),
          limit: integer(),
          type: String.t(),
          extra: integer(),
          reset: String.t()
        }

  @type traffic_details_day :: %{
          host: map(),
          bytes: integer()
        }

  @doc """
  Get traffic information for each hoster.

  ## Returns

    - `{:ok, traffic}` - Map of host to traffic info
    - `{:error, reason}` on failure
  """
  @spec get(Client.t()) :: {:ok, map()} | {:error, term()}
  def get(%Client{} = client) do
    case Client.get(client, "/traffic") do
      {:ok, body, _headers} ->
        traffic =
          body
          |> Enum.map(fn {key, value} ->
            {key,
             %{
               left: value["left"],
               bytes: value["bytes"],
               links: value["links"],
               limit: value["limit"],
               type: value["type"],
               extra: value["extra"],
               reset: value["reset"]
             }}
          end)
          |> Map.new()

        {:ok, traffic}

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Get detailed traffic information.

  ## Parameters

    - `client` - The Real-Debrid client
    - `opts` - Optional keyword list:
      - `:start` - Start date (ISO format)
      - `:end` - End date (ISO format)

  ## Returns

    - `{:ok, details}` - Map of date to traffic details
    - `{:error, reason}` on failure
  """
  @spec get_details(Client.t(), keyword()) :: {:ok, map()} | {:error, term()}
  def get_details(%Client{} = client, opts \\ []) do
    params =
      []
      |> maybe_add_param(:start, Keyword.get(opts, :start))
      |> maybe_add_param(:end, Keyword.get(opts, :end))
      |> Map.new()

    case Client.get(client, "/traffic/details", params: params) do
      {:ok, body, _headers} ->
        details =
          body
          |> Enum.map(fn {key, value} ->
            {key,
             %{
               host: value["host"] || %{},
               bytes: value["bytes"]
             }}
          end)
          |> Map.new()

        {:ok, details}

      {:error, reason} ->
        {:error, reason}
    end
  end

  defp maybe_add_param(params, _key, nil), do: params
  defp maybe_add_param(params, key, value), do: [{key, value} | params]
end
