defmodule RealDebrid.Api.Auth do
  @moduledoc """
  Authentication-related API functions.
  """

  alias RealDebrid.Client

  @doc """
  Disables the current access token.

  ## Returns

    - `:ok` on success (204 response)
    - `{:error, reason}` on failure
  """
  @spec disable_access_token(Client.t()) :: :ok | {:error, term()}
  def disable_access_token(%Client{} = client) do
    case Client.get(client, "/disable_access_token") do
      {:ok, _body, _headers} ->
        :ok

      {:error, reason} ->
        {:error, reason}
    end
  end
end
