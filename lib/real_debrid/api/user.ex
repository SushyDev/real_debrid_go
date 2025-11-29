defmodule RealDebrid.Api.User do
  @moduledoc """
  User information API functions.
  """

  alias RealDebrid.Client

  @type user :: %{
          id: integer(),
          username: String.t(),
          email: String.t(),
          points: integer(),
          locale: String.t() | nil,
          avatar: String.t(),
          type: String.t(),
          premium: integer(),
          expiration: String.t()
        }

  @doc """
  Get current user information.

  ## Returns

    - `{:ok, user}` on success
    - `{:error, reason}` on failure
  """
  @spec get(Client.t()) :: {:ok, user()} | {:error, term()}
  def get(%Client{} = client) do
    case Client.get(client, "/user") do
      {:ok, body, _headers} ->
        {:ok,
         %{
           id: body["id"],
           username: body["username"],
           email: body["email"],
           points: body["points"],
           locale: body["locale"],
           avatar: body["avatar"],
           type: body["type"],
           premium: body["premium"],
           expiration: body["expiration"]
         }}

      {:error, reason} ->
        {:error, reason}
    end
  end
end
