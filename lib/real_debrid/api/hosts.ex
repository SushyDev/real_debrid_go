defmodule RealDebrid.Api.Hosts do
  @moduledoc """
  Hosts-related API functions.
  """

  alias RealDebrid.Client

  @type host_basic :: %{
          id: String.t(),
          name: String.t(),
          image: String.t()
        }

  @type host_status :: %{
          id: String.t(),
          name: String.t(),
          image: String.t(),
          supported: integer(),
          status: String.t(),
          check_time: String.t(),
          competitors_status: map()
        }

  @doc """
  Get list of supported hosts.

  ## Returns

    - `{:ok, hosts}` - Map of host ID to host info
    - `{:error, reason}` on failure
  """
  @spec get(Client.t()) :: {:ok, map()} | {:error, term()}
  def get(%Client{} = client) do
    case Client.get(client, "/hosts") do
      {:ok, body, _headers} ->
        hosts =
          body
          |> Enum.map(fn {key, value} ->
            {key,
             %{
               id: value["id"],
               name: value["name"],
               image: value["image"]
             }}
          end)
          |> Map.new()

        {:ok, hosts}

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Get status of supported hosts.

  ## Returns

    - `{:ok, hosts}` - Map of host ID to host status
    - `{:error, reason}` on failure
  """
  @spec get_status(Client.t()) :: {:ok, map()} | {:error, term()}
  def get_status(%Client{} = client) do
    case Client.get(client, "/hosts/status") do
      {:ok, body, _headers} ->
        hosts =
          body
          |> Enum.map(fn {key, value} ->
            {key,
             %{
               id: value["id"],
               name: value["name"],
               image: value["image"],
               supported: value["supported"],
               status: value["status"],
               check_time: value["check_time"],
               competitors_status: value["competitors_status"] || %{}
             }}
          end)
          |> Map.new()

        {:ok, hosts}

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Get supported hosts regex patterns.

  ## Returns

    - `{:ok, patterns}` - List of regex patterns
    - `{:error, reason}` on failure
  """
  @spec get_regex(Client.t()) :: {:ok, [String.t()]} | {:error, term()}
  def get_regex(%Client{} = client) do
    case Client.get(client, "/hosts/regex") do
      {:ok, body, _headers} -> {:ok, body}
      {:error, reason} -> {:error, reason}
    end
  end

  @doc """
  Get supported hosts folder regex patterns.

  ## Returns

    - `{:ok, patterns}` - List of folder regex patterns
    - `{:error, reason}` on failure
  """
  @spec get_regex_folder(Client.t()) :: {:ok, [String.t()]} | {:error, term()}
  def get_regex_folder(%Client{} = client) do
    case Client.get(client, "/hosts/regexFolder") do
      {:ok, body, _headers} -> {:ok, body}
      {:error, reason} -> {:error, reason}
    end
  end

  @doc """
  Get supported hosts domains.

  ## Returns

    - `{:ok, domains}` - List of supported domains
    - `{:error, reason}` on failure
  """
  @spec get_domains(Client.t()) :: {:ok, [String.t()]} | {:error, term()}
  def get_domains(%Client{} = client) do
    case Client.get(client, "/hosts/domains") do
      {:ok, body, _headers} -> {:ok, body}
      {:error, reason} -> {:error, reason}
    end
  end
end
