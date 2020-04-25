USE [GoKube]
GO

/****** Object:  StoredProcedure [Customer].[usp_CustomerCredentials_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_CustomerCredentials_sel')
BEGIN
    DROP PROCEDURE [Customer].[usp_CustomerCredentials_sel]
END
GO

/****** Object:  StoredProcedure [Customer].[usp_CustomerCredentials_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-04-23  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [Customer].[usp_CustomerCredentials_sel]
    (
    @CustomerID int
)
AS
BEGIN

    SET NOCOUNT ON;

    SELECT [ClientAPIKey]
      , [Secret]
      , [CustomerID]
      , [StartDate]
      , [EndDate]
      , [RowInserted]
    FROM [GoKube].[Customer].[CustomerCredentials]
    WHERE 
		CustomerID = @CustomerID
        AND EndDate IS NULL

END

GO